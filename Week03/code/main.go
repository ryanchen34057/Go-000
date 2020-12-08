package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func serveApp(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	fmt.Println("Start main server on 8080")
	server.ListenAndServe()
	select {
	case <-ctx.Done():
		fmt.Println("Main server exit")
		return ctx.Err()
	}
}

func serveDebug(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8081",
		Handler: nil,
	}
	fmt.Println("Start debug server on 8081")
	server.ListenAndServe()
	select {
	case <-ctx.Done():
		fmt.Println("Debug server exit")
		return ctx.Err()
	}
}

func processSignals(sigs chan os.Signal, cancel context.CancelFunc) {
	<-sigs
	cancel()
}

func main() {
	g, errctx := errgroup.WithContext(context.Background())
	ctx, cancel := context.WithCancel(errctx)

	// serveApp
	g.Go(func() error {
		return serveApp(ctx)
	})

	// serve debug
	g.Go(func() error {
		return serveDebug(ctx)
	})

	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go processSignals(signals, cancel)

	if err := g.Wait(); err == nil {
		fmt.Println("Program exit")
	}
}
