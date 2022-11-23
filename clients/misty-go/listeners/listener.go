package listeners

// MistyListener interface specifies the required methods need by a
// type to act as a misty listener client.
type MistyListener interface {
	// Connect method sends a PUT request to the misty broker to
	// initiate a connection
	Connect(host, port string) error
	// Close method sends a DELETE request to the misty broker to
	// remove the client from the broker list
	Close() error
	// Subscribe method connects the listener to a topic
	Subscribe(topic string) error
}
