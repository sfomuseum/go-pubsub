package subscribe

import (
	"context"
	"flag"
	"fmt"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	SubscriberURI string
	PublisherURI  string
}

func OptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "PUBSUB")

	if err != nil {
		return nil, fmt.Errorf("Failed to derive flags from environment variables, %w", err)
	}

	opts := &RunOptions{
		SubscriberURI: subscriber_uri,
		PublisherURI:  publisher_uri,
	}

	return opts, nil
}
