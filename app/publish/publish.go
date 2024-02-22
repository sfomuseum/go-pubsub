package publish

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"

	"github.com/sfomuseum/go-pubsub/publisher"
)

func Run(ctx context.Context, logger *slog.Logger) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs, logger)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet, logger *slog.Logger) error {

	opts, err := OptionsFromFlagSet(ctx, fs)

	if err != nil {
		return fmt.Errorf("Failed to derive options from flagset, %w", err)
	}

	return RunWithOptions(ctx, opts, logger)
}

func RunWithOptions(ctx context.Context, opts *RunOptions, logger *slog.Logger) error {

	pub, err := publisher.NewPublisher(ctx, opts.PublisherURI)

	if err != nil {
		return fmt.Errorf("Failed to create new publisher, %v", err)
	}

	defer pub.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		pub.Publish(ctx, scanner.Text())
	}

	if scanner.Err() != nil {
		return fmt.Errorf("Failed to scan")
	}

	return nil
}
