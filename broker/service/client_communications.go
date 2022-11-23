package service

import (
	"fmt"
	"log"
	"net/http"
	utils "rohitsingh/misty-utils"
)

// Broadcast sends a published message to all subscribed clients
func Broadcast(topic string, message string) error {
	// Get all the listeners associated with the topic from the repository
	listeners, err := repo.GetTopicListeners(topic)
	if err != nil {
		return err
	}
	for _, l := range listeners {
		// Send a PUT request to each listener
		httpUrl := fmt.Sprintf("%s/%s", l, topic)
		log.Printf("Broadcasting %s to %s", message, httpUrl)
		if err := utils.PutString(httpUrl, message, http.StatusAccepted); err != nil {
			return err
		}

		// req, err := http.NewRequest(
		// 	http.MethodPut,
		// 	httpUrl,
		// 	bytes.NewBuffer([]byte(message)),
		// )
		// if err != nil {
		// 	return fmt.Errorf("error while creating PUT request: %q", err)
		// }
		// req.Header.Set("Content-Type", "text/plain")
		// resp, err := http.DefaultClient.Do(req)
		// if err != nil {
		// 	return fmt.Errorf("error while sending PUT request: %q", err)
		// }
		// if resp.StatusCode != http.StatusAccepted {
		// 	return fmt.Errorf("error while making PUT request: %s", http.StatusText(resp.StatusCode))
		// }
	}
	return nil
}

// CloseConnections sends a PUT message to all subscribed clients
// to let them know that the server is shutting down
func CloseConnections() error {

	return nil
}
