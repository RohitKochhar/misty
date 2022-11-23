package actions

import (
	"fmt"
	"net/http"
)

// Publish sends a PUT request to send a given
// message to a given topic and returns the
// response error status
func Publish(host string, port int, topic string, message string) error {
	// PUT the message on the corresponding topic
	// ToDo: Add topic sanitizing / function to create HTTP URL with rules
	httpUrl := fmt.Sprintf("http://%s:%d/topic%s", host, port, topic)
	if err := PutString(httpUrl, message, http.StatusOK); err != nil {
		return err
	}
	fmt.Printf("[PUBLISH] %s --> %s\n", message, httpUrl)
	return nil
}
