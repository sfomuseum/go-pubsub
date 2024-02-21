package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/sfomuseum/go-pubsub/subscriber"
)

func main() {

	var subscriber_uri string

	flag.StringVar(&subscriber_uri, "subscriber-uri", "", "...")

	flag.Parse()

	ctx := context.Background()

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	sub, err := subscriber.NewSubscriber(ctx, subscriber_uri)

	if err != nil {
		log.Fatalf("Failed to create new subscriber, %v", err)
	}

	defer sub.Close()
	
	msg_ch := make(chan string)
	done_ch := make(chan bool)

	go func() {

		for {
			select {
			case <-ctx.Done():
				return
			case <-done_ch:
				return
			case msg := <-msg_ch:
				fmt.Println(msg)
			default:
				//
			}
		}
	}()

	log.Println("Listening for messages")
	err = sub.Listen(ctx, msg_ch)

	done_ch <- true

	if err != nil {
		log.Fatalf("Failed to listen, %v", err)
	}
}
