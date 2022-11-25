package main

import (
	"fmt"
	"log"
	"os"
	"rohitsingh/misty-go/clients"
	"strconv"
)

/*
- Sample service reverses the string and appends
	the number of the service to the string and sends it back
*/

func main() {
	// The number of the client is set through env
	clientHost := fmt.Sprintf("service-%s", os.Getenv("SERVICE_NUMBER"))
	// assign a unique port
	inc, err := strconv.Atoi(os.Getenv("SERVICE_NUMBER"))
	if err != nil {
		log.Panic(err)
		os.Exit(1)
	}
	clientPort := 2397 + inc
	// Create a new listener to connect to the broker
	l := clients.NewListener(clientHost, clientPort)
	// Connect to the broker
	if err := l.Connect("broker", 1315); err != nil {
		log.Fatal(err)
	}
	// Subscribe to messages sent to this client's topic
	if err := l.Subscribe(clientHost); err != nil {
		log.Fatal(err)
	}
	// Listen for messages
	if err := l.Listen(); err != nil {
		log.Fatal(err)
	}
}
