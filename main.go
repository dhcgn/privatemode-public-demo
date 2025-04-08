package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/edgelesssys/continuum/privatemode-proxy/cmd"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		cancel()
	}()

	defer cancel()

	cmd := cmd.New()
	if err := cmd.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
