package actions

import (
	"fmt"
	"log"
	"net/http"
	"rohitsingh/misty-cli/utils"
)

// Publish sends a PUT request to send a given
// message to a given topic and returns the
// response error status
func Publish(host string, port int, topic string, message string) error {
	// Sanitize the topic
	topic, err := utils.SanitizeTopic(topic)
	if err != nil {
		return err
	}
	// PUT the message on the corresponding topic
	httpUrl := fmt.Sprintf("http://%s:%d/topic%s", host, port, topic)
	if err := utils.PutString(httpUrl, message, http.StatusOK); err != nil {
		return err
	}
	log.Printf("[PUBLISH] %s --> %s\n", message, httpUrl)
	return nil
}
