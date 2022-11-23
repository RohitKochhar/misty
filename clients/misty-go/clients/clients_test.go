package clients_test

import (
	"fmt"
	"log"
	"rohitsingh/misty-go/clients"
	"testing"
)

// Note: These tests will fail unless a test
// misty broker is up and running on localhost:1315

// ConnectListener is a helper function that
// creates a Listener and connects it
// to a misty server running on localhost:1315
func ConnectListener(t *testing.T) *clients.Listener {
	// Mark this as a helper since it should be
	// run at the beginning of all tests
	t.Helper()
	// Define the parameters of the running misty broker
	brokerName, brokerPort := "localhost", 1315
	// Create a new blocking listener object with default params
	Listener := clients.NewListener("", 0)
	// Send a connection request to the running misty broker
	if err := Listener.Connect(brokerName, brokerPort); err != nil {
		t.Fatalf("error while trying to connect to misty broker at %s:%d : %q", brokerName, brokerPort, err)
	}
	return Listener
}

// TestClose tests that a Listener can
// close its connection to a misty broker
func TestClose(t *testing.T) {
	// Connect to the misty server
	l := ConnectListener(t)
	// and disconnect from it!
	defer l.Close()
}

// TestSubscribe tests that a Listener can
// subscribe to a topic successfully
func TestSubscribe(t *testing.T) {
	// Connect to the misty server
	l := ConnectListener(t)
	defer l.Close()
	// Subscribe to an incorrectly formatted sample topic
	topic := "sample-topic/"
	if err := l.Subscribe(topic); err != nil {
		t.Fatalf("error while subscribing to %s: %q", topic, err)
	}
}

// TestClientIntegration starts a listener, subscribes to topics
// and prints messages until the test is cancelled
func TestClientIntegration(t *testing.T) {
	// Connect to the misty server
	l := ConnectListener(t)
	defer l.Close()
	// Subscribe to a few topics
	topics := []string{"/test1", "/test2/", "test3/", "another-topic"}
	for _, topic := range topics {
		if err := l.Subscribe(topic); err != nil {
			t.Fatalf("error while subscribing to %s: %q", topic, err)
		}
	}
	// Create a publisher
	p := clients.NewPublisher()
	// Listen to for messages concurrently as we send messages
	errCh := make(chan error)
	doneCh := make(chan bool)
	go func() {
		if err := l.Listen(); err != nil {
			errCh <- fmt.Errorf("error while listening for messages: %q", err)
		}
	}()
	go func() {
		for _, topic := range topics {
			for i := 0; i < 1000; i++ {
				p.Publish("localhost", 1315, topic, fmt.Sprintf("message%d", i))
			}
		}
		doneCh <- true
	}()
	select {
	case err := <-errCh:
		log.Println(err)
		if err := l.Close(); err != nil {
			t.Fatalf("couldn't remove listener from broker list due to error: %q", err)
		}
	case <-doneCh:
		l.Close()
		return
	}
}
