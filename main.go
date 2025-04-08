package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/edgelesssys/continuum/privatemode-proxy/cmd"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file from the current directory
	err := godotenv.Load()
	if err != nil {
		os.Stderr.WriteString("Error loading .env file: " + err.Error() + "\n")
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		cancel()
	}()

	defer cancel()

	command := cmd.New()

	// Set the API key from an environment variable
	apiKey := os.Getenv("PRIVATEMODE_API_KEY")
	if apiKey != "" {
		if err := command.Flags().Set("apiKey", apiKey); err != nil {
			// Handle error if the flag doesn't exist or can't be set
			os.Stderr.WriteString("Error setting API key from environment variable: " + err.Error() + "\n")
			os.Exit(1)
		}
	}

	if err := command.ExecuteContext(ctx); err != nil {
		os.Stderr.WriteString("Error executing command: " + err.Error() + "\n")
		os.Exit(1)
	}
}
