package clients

// MistyListener interface specifies the required methods need by a
// type to act as a misty listener client.
type MistyListener interface {
	// Connect method sends a PUT request to the misty broker to
	// initiate a connection
	Connect(host string, port int) error
	// Close method sends a DELETE request to the misty broker to
	// remove the client from the broker list
	Close() error
	// Subscribe method connects the listener to a topic
	Subscribe(topic string) error
}

// MistyPublisher interface specifies the required methods needed by
// a type to act as a misty publisher client
type MistyPublisher interface {
	// Publish sends a PUT request to the server containing some data
	Publish(host string, port int, topic string, message string) error
}
