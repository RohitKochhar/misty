package main

import (
	"log"
	"rohitsingh/misty-go/clients"
)

/*
- Entrypoint listens for PUT requests on an unexposed IP :2396 (only exposed to docker)
  and prints the incoming logs to stdout
*/

func main() {
	// Create a listener to connect to broker
	l := clients.NewListener("logger", 2396)
	defer l.Close()
	// Connect to the broker
	if err := l.Connect("broker", 1315); err != nil {
		log.Fatal(err)
	}
	// Subscribe to log topic
	if err := l.Subscribe("/logs"); err != nil {
		log.Fatal(err)
	}
	// Listen for messages
	if err := l.Listen(); err != nil {
		log.Fatal(err)
	}
}
