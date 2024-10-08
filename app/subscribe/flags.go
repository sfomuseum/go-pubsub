package subscribe

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var subscriber_uri string
var publisher_uri string

func DefaultFlagSet() *flag.FlagSet {
	fs := flagset.NewFlagSet("subscribe")
	fs.StringVar(&subscriber_uri, "subscriber-uri", "", "A valid sfomuseum/go-pubsub/subscriber.Subscriber URI")
	fs.StringVar(&publisher_uri, "publisher-uri", "stdout://?newline=true", "A valid sfomuseum/go-pubsub/publisher.Publisher URI")
	return fs
}
