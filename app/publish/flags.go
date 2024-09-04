package publish

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var publisher_uri string
var mode string

func DefaultFlagSet() *flag.FlagSet {
	fs := flagset.NewFlagSet("publish")
	fs.StringVar(&publisher_uri, "publisher-uri", "", "A valid sfomuseum/go-pubsub/publisher.Publisher URI")
	fs.StringVar(&mode, "mode", "", "Optional flag to signal whether data should be read for an alternate source. Valid options are: readline, stdin.")
	return fs
}
