package service

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// rootHandler handles requests made to the root of the server
func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	host := viper.Get("host")
	port := viper.Get("port")
	replyContent := fmt.Sprintf("misty server ready to accept connections at http://%s:%d\n", host, port)
	replyTextContent(w, r, http.StatusOK, replyContent)
}

// publishHandler handles requests made to a single-level topic
func publishHandler(w http.ResponseWriter, r *http.Request) {
	// Get the topic and value from the request
	topic := mux.Vars(r)["topic"]
	message, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, "Could not read request body")
		return
	}
	log.Printf("Received %s on %s", string(message), topic)
	// Broadcast the message to subscribed clients
	replyTextContent(w, r, http.StatusOK, "")
	Broadcast(topic, string(message))
}

// addListenerHandler handles requests to add a listener to a topic
func addListenerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the topic that the listener wants to subscribe to
	topic := mux.Vars(r)["topic"]
	// Get the information about the listener from the packet body
	listener, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, "Could not add new listener")
		return
	}
	log.Printf("Adding new listener %s to topic=%s", string(listener), topic)
	// Add the listener to the repository
	if err := repo.NewListener(string(listener)); err != nil {
		log.Printf("error while adding listener (%s): %q", string(listener), err)
	} else {
		// If adding the listener was successful, add the topic to it
		if err := repo.AddTopicToListener(string(listener), topic); err != nil {
			log.Printf("error while adding topic (%s) to listener (%s): %q", topic, string(listener), err)
		}
		replyTextContent(w, r, http.StatusAccepted, "accepted new listener")
	}
	log.Printf("Successfully added listener %s to topic=%s", string(listener), topic)
}
