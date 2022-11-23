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
	// If _ has been provided as the topic, don't add a topic to the listener
	isConnectionRequest := topic == "_"
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
		// Check if we are handling a basic connection request or a connect and subscribe
		if !isConnectionRequest {
			// If adding the listener was successful, add the topic to it
			if err := repo.AddTopicToListener(string(listener), topic); err != nil {
				log.Printf("error while adding topic (%s) to listener (%s): %q", topic, string(listener), err)
			}
			replyTextContent(w, r, http.StatusAccepted, "accepted new listener")
			log.Printf("Successfully added listener %s to topic=%s", string(listener), topic)
		} else {
			replyTextContent(w, r, http.StatusAccepted, "accepted new listener")
			log.Printf("Successfully added listener %s to broker list", string(listener))
		}
	}

}

func deleteListenerHandler(w http.ResponseWriter, r *http.Request) {
	// Get the listener address from the body
	listenerAddr, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, "Could not delete listener")
		return
	}
	listener := string(listenerAddr)
	log.Printf("Received request to delete listener: %s", listener)
	// Delete the listener from the repository
	if err := repo.DeleteListener(listener); err != nil {
		replyError(w, r, http.StatusInternalServerError, "Could not delete listener")
		return
	}
	replyTextContent(w, r, http.StatusOK, fmt.Sprintf("deleted listener %s", listener))
	log.Printf("Successfully deleted listener: %s", listener)
}

// listenerSubscribeHandler handles requests to add a topic to a listener's subscription list
func listenerSubscribeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the topic that the listener wants to subscribe to
	topic := mux.Vars(r)["topic"]
	// Get the information about the listener from the packet body
	listener, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, "Could not add new listener")
		return
	}
	if err := repo.AddTopicToListener(string(listener), topic); err != nil {
		log.Printf("error while subscribing listener (%s) to topic (%s): %q", string(listener), topic, err)
	}
	replyTextContent(w, r, http.StatusAccepted, "subscribed listener to topic")
	log.Printf("Subscribing listener %s to topic=%s", string(listener), topic)
}
