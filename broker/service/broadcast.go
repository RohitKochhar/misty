package service

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// Broadcast sends a published message to all subscribed clients
func Broadcast(topic string, message string) error {
	// Get all the listeners associated with the topic from the repository
	listeners, err := repo.GetTopicListeners(topic)
	if err != nil {
		return err
	}
	for _, l := range listeners {
		httpUrl := fmt.Sprintf("%s/%s", l, topic)
		log.Printf("Broadcasting %s to %s", message, httpUrl)
		req, err := http.NewRequest(
			http.MethodPut,
			httpUrl,
			bytes.NewBuffer([]byte(message)),
		)
		if err != nil {
			return fmt.Errorf("error while creating PUT request: %q", err)
		}
		req.Header.Set("Content-Type", "text/plain")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return fmt.Errorf("error while sending PUT request: %q", err)
		}
		if resp.StatusCode != http.StatusAccepted {
			return fmt.Errorf("error while making PUT request: %s", http.StatusText(resp.StatusCode))
		}
	}
	return nil
}
