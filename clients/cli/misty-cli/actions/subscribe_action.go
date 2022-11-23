package actions

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// SubscribeAction creates a server to accept messages by the broker
func SubscribeAction(brokerHost string, brokerPort int, topic string) error {
	// Define the parameters used by the listener as a HTTP server
	// ToDo: Dynamically get configuration
	listenerHost := "localhost"
	listenerPort := 1111
	// Send a PUT request to the server to
	// let it know we are interested in the topic
	log.Printf("Sending subscribe request to broker at %s:%d...\n", brokerHost, brokerPort)
	if err := requestSubscribe(brokerHost, brokerPort, listenerHost, listenerPort, topic); err != nil {
		log.Panicf("Subscribe result unsuccessful: %q\n", err)
		return err
	}
	log.Println("Subscibe result successful!")
	// Create a HTTP server that only listens on the given topic
	r := mux.NewRouter()
	r.HandleFunc(fmt.Sprintf("/%s", topic), subscribeHandler).Methods(http.MethodPut)
	log.Printf("listening for messages on %s...\n", topic)
	// ToDo: Add context handling for graceful shutdown
	if err := http.ListenAndServe(fmt.Sprintf(":%d", listenerPort), r); err != nil {
		return err
	}
	return nil
}

// subscribeHandler handles incoming PUT requests to a topic
func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value from the request
	message, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, "could not accept packet")
	}
	replyTextContent(w, r, http.StatusAccepted, "accepted packet")
	log.Printf("[RECEIVED ON %s]: %s", r.URL.Path, string(message))
}

// requestSubscribe sends a message to the broker letting it know
// that this client wants to receive messages for a given topic
func requestSubscribe(brokerHost string, brokerPort int, listenerHost string, listenerPort int, topic string) error {
	// Send a PUT request to the broker containing listener server information
	httpUrl := fmt.Sprintf("http://%s:%d/listeners/%s/add", brokerHost, brokerPort, topic)
	message := fmt.Sprintf("http://%s:%d", listenerHost, listenerPort)
	if err := PutString(httpUrl, message, http.StatusAccepted); err != nil {
		return err
	}
	return nil
}
