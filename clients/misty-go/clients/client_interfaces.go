package clients

import "fmt"

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

// SanitizeTopic parses a topic string to ensure that
// all topics are formatted uniformly with the form"
// "/{TOPIC}/{SUBTOPIC}/{SUBSUBTOPIC}" (leading /, no trailing /)
func SanitizeTopic(topic string) (string, error) {
	var sanitizedTopic string
	// Check if the last character is a dash
	if topic[len(topic)-1] == '/' {
		sanitizedTopic = topic[:len(topic)-1]
	} else {
		sanitizedTopic = topic
	}
	// Check if the first character in the topic is a /
	if topic[0] != '/' {
		sanitizedTopic = fmt.Sprintf("/%s", sanitizedTopic)
	}

	return sanitizedTopic, nil
}
