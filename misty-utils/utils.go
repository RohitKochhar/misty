package utils

import "fmt"

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
