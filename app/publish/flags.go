package publish

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var publisher_uri string
var mode string

func DefaultFlagSet() *flag.FlagSet {
	fs := flagset.NewFlagSet("publish")
	fs.StringVar(&publisher_uri, "publisher-uri", "", "...")
	fs.StringVar(&mode, "mode", "", "...")
	return fs
}
