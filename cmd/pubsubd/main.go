package main

import (
	"flag"
	"log"

	"github.com/sfomuseum/go-pubsub/redis"
)

func main() {

	var host string
	var port int
	var debug bool

	flag.StringVar(&host, "host", "localhost", "The hostname to listen on.")
	flag.IntVar(&port, "port", 6379, "The port number to listen on.")
	flag.BoolVar(&debug, "debug", false, "Print all RESP commands to STDOUT.")

	flag.Parse()

	server, err := redis.NewPubSubServer(host, port)

	if err != nil {
		log.Fatalf("Failed to create new server, %v", err)
	}

	server.Debug = debug

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to serve requests, %v", err)
	}
}
