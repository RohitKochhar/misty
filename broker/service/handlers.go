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
		replyError(w, r, http.StatusInternalServerError, "Could not ready request body")
		return
	}
	log.Printf("Received %s on %s", string(message), topic)
	replyTextContent(w, r, http.StatusOK, "")
}
