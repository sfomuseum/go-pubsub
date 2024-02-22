package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/sfomuseum/go-pubsub/app/publish"
)

func main() {

	ctx := context.Background()
	logger := slog.Default()

	err := publish.Run(ctx, logger)

	if err != nil {
		logger.Error("Failed to run publish tool", "error", err)
		os.Exit(1)
	}
}
