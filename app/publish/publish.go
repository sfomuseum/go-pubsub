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

	switch opts.Mode {
	case "stdin":

		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			pub.Publish(ctx, scanner.Text())
		}

		if scanner.Err() != nil {
			return fmt.Errorf("Failed to scan")
		}

	case "readline":

		logger.Info("Type messages to publish")

		for {

			// Please fix me: This treats spaces in the input
			// as separate messages...

			var input string
			fmt.Scanln(&input)

			if input == "." {
				break
			}

			err := pub.Publish(ctx, input)

			if err != nil {
				return fmt.Errorf("Failed to publish message, %w", err)
			}

		}

	default:

		err := pub.Publish(ctx, opts.Message)

		if err != nil {
			return fmt.Errorf("Failed to publish message, %w", err)
		}
	}

	return nil
}
