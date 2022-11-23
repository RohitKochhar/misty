package listeners_test

import (
	"rohitsingh/misty-go/listeners"
	"testing"
)

// Note: These tests will fail unless a test
// misty broker is up and running on localhost:1315

// ConnectListener is a helper function that
// creates a Listener and connects it
// to a misty server running on localhost:1315
func ConnectListener(t *testing.T) *listeners.Listener {
	// Mark this as a helper since it should be
	// run at the beginning of all tests
	t.Helper()
	// Define the parameters of the running misty broker
	brokerName, brokerPort := "localhost", 1315
	// Create a new blocking listener object with default params
	Listener := listeners.NewListener("", 0)
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

// TestListenerIntegration starts a listener, subscribes to topics
// and prints messages until the test is cancelled
func TestListenerIntegration(t *testing.T) {
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
	// Listen to for messages
	if err := l.Listen(); err != nil {
		t.Fatalf("error while listening for messages: %q", err)
	}
}
