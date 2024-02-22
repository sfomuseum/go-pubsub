package publish

import (
	"context"
	"flag"
	"fmt"
	"strings"

	"github.com/sfomuseum/go-flags/flagset"
)

type RunOptions struct {
	PublisherURI string
	Mode         string
	Message      string
}

func OptionsFromFlagSet(ctx context.Context, fs *flag.FlagSet) (*RunOptions, error) {

	flagset.Parse(fs)

	err := flagset.SetFlagsFromEnvVars(fs, "PUBSUB")

	if err != nil {
		return nil, fmt.Errorf("Failed to derive flags from environment variables, %w", err)
	}

	opts := &RunOptions{
		PublisherURI: publisher_uri,
		Mode:         mode,
	}

	switch mode {
	case "stdin", "readline":
		// pass
	default:
		opts.Message = strings.Join(fs.Args(), " ")
	}

	return opts, nil
}
