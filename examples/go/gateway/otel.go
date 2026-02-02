package main

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/log/global"
	"go.opentelemetry.io/otel/propagation"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const schemaName = "https://github.com/grafana/docker-otel-lgtm"

var (
	tracer = otel.Tracer(schemaName)
	logger = otelslog.NewLogger(schemaName)
	meter  = otel.Meter(schemaName)
)

func setupOTelSDK(ctx context.Context) (shutdown func(context.Context) error, err error) {
	var shutdownFuncs []func(context.Context) error

	shutdown = func(ctx context.Context) error {
		var e error
		for _, fn := range shutdownFuncs {
			e = errors.Join(e, fn(ctx))
		}
		shutdownFuncs = nil
		return e
	}

	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}

	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	traceExporter, err := otlptrace.New(ctx, otlptracehttp.NewClient())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(traceExporter))
	shutdownFuncs = append(shutdownFuncs, tp.Shutdown)
	otel.SetTracerProvider(tp)

	metricExporter, err := otlpmetrichttp.New(ctx)
	if err != nil {
		return nil, err
	}
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(metricExporter)),
	)
	shutdownFuncs = append(shutdownFuncs, mp.Shutdown)
	otel.SetMeterProvider(mp)

	logExporter, err := otlploghttp.New(ctx, otlploghttp.WithInsecure())
	if err != nil {
		return nil, err
	}
	lp := sdklog.NewLoggerProvider(sdklog.WithProcessor(sdklog.NewBatchProcessor(logExporter)))
	shutdownFuncs = append(shutdownFuncs, lp.Shutdown)
	global.SetLoggerProvider(lp)

	if e := runtime.Start(runtime.WithMinimumReadMemStatsInterval(time.Second)); e != nil {
		logger.ErrorContext(ctx, "otel runtime instrumentation failed", slog.Any("error", e))
	}

	return shutdown, nil
}
