package main

import (
	"bufio"
	"context"
	"flag"
	"log"
	"os"

	"github.com/sfomuseum/go-pubsub/publisher"
)

func main() {

	var publisher_uri string

	flag.StringVar(&publisher_uri, "publisher-uri", "", "...")

	flag.Parse()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pub, err := publisher.NewPublisher(ctx, publisher_uri)

	if err != nil {
		log.Fatalf("Failed to create new publisher, %v", err)
	}

	defer pub.Close()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		pub.Publish(ctx, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatalf("Failed to scan")
	}
}
