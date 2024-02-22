package subscribe

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/sfomuseum/go-pubsub/publisher"
	"github.com/sfomuseum/go-pubsub/subscriber"
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

	sub, err := subscriber.NewSubscriber(ctx, opts.SubscriberURI)

	if err != nil {
		return fmt.Errorf("Failed to create new subscriber, %v", err)
	}

	defer sub.Close()

	pub, err := publisher.NewPublisher(ctx, opts.PublisherURI)

	if err != nil {
		return fmt.Errorf("Failed to create new publisher, %v", err)
	}

	defer pub.Close()

	msg_ch := make(chan string)
	done_ch := make(chan bool)

	publish := func(ctx context.Context, msg string) {

		err := pub.Publish(ctx, msg)

		if err != nil {
			logger.Error("Failed to publish message", "error", err)
		}
	}

	go func() {

		for {
			select {
			case <-ctx.Done():
				return
			case <-done_ch:
				return
			case msg := <-msg_ch:
				go publish(ctx, msg)
			default:
				//
			}
		}
	}()

	logger.Info("Listening for messages")
	err = sub.Listen(ctx, msg_ch)

	done_ch <- true

	if err != nil {
		return fmt.Errorf("Failed to listen, %v", err)
	}

	return nil
}
