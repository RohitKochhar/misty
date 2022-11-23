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
	}
	return nil
}

// CloseConnections sends a PUT message to all subscribed clients
// to let them know that the server is shutting down
func CloseConnections() error {
	// Get all the listeners associated with the broker
	listeners, err := repo.GetAllListeners()
	if err != nil {
		return err
	}
	for _, l := range listeners {
		// Send a PUT request to each listener with an empty
		httpUrl := fmt.Sprintf("%s/broker/down", l)
		log.Printf("Letting %s know broker is going down", l)
		// Let the client know we are going down, but we can't
		// handle errors (since we are going down)
		_ = utils.PutString(httpUrl, "server going down!", http.StatusAccepted)
	}
	return nil
}
