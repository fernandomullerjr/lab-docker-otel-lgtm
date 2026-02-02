package main

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	downstreamLatencyMs = mustHistogram("gateway.downstream.latency_ms", "Downstream rolldice call latency in milliseconds")
	downstreamErrors    = mustCounter("gateway.downstream.errors_total", "Downstream rolldice call errors")
)

func mustHistogram(name, desc string) metric.Int64Histogram {
	h, err := meter.Int64Histogram(name, metric.WithDescription(desc))
	if err != nil {
		panic(err)
	}
	return h
}

func mustCounter(name, desc string) metric.Int64Counter {
	c, err := meter.Int64Counter(name, metric.WithDescription(desc))
	if err != nil {
		panic(err)
	}
	return c
}

func rollViaDownstream(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "gateway.roll")
	defer span.End()

	baseURL := os.Getenv("ROLLDICE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:8081"
	}
	url := baseURL + "/rolldice"

	client := &http.Client{
		Timeout:   3 * time.Second,
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		downstreamErrors.Add(ctx, 1, metric.WithAttributes(attribute.String("error.type", "new_request")))
		http.Error(w, "failed to create request", http.StatusInternalServerError)
		return
	}

	resp, err := client.Do(req)
	latencyMs := time.Since(start).Milliseconds()
	downstreamLatencyMs.Record(ctx, latencyMs, metric.WithAttributes(
		attribute.String("downstream.service", "rolldice"),
		attribute.Int("http.status_code", func() int {
			if resp == nil {
				return 0
			}
			return resp.StatusCode
		}()),
	))

	if err != nil {
		span.RecordError(err)
		downstreamErrors.Add(ctx, 1, metric.WithAttributes(attribute.String("error.type", "do_request")))
		logger.ErrorContext(ctx, "downstream request failed", slog.Any("error", err), slog.String("url", url))
		http.Error(w, "downstream error", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	logger.InfoContext(ctx, "downstream response",
		slog.String("url", url),
		slog.Int("status", resp.StatusCode),
		slog.Int64("latency_ms", latencyMs),
		slog.String("result", string(body)),
	)

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		downstreamErrors.Add(ctx, 1, metric.WithAttributes(attribute.String("error.type", "http_status")))
		http.Error(w, "downstream returned non-2xx", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprintf(w, "gateway_ok latency_ms=%d result=%s", latencyMs, body)
}
