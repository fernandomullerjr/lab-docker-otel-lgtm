package main

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() (err error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		return
	}
	defer func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}()

	mux := http.NewServeMux()
	mux.Handle("/roll", otelhttp.WithRouteTag("/roll", http.HandlerFunc(rollViaDownstream)))
	mux.Handle("/healthz", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) }))

	srv := &http.Server{
		Addr:         ":8080",
		BaseContext:  func(_ net.Listener) context.Context { return ctx },
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      otelhttp.NewHandler(mux, "/"),
	}

	srvErr := make(chan error, 1)
	go func() { srvErr <- srv.ListenAndServe() }()

	select {
	case err = <-srvErr:
		return
	case <-ctx.Done():
		stop()
	}

	err = srv.Shutdown(context.Background())
	return
}
