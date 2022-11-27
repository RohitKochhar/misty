package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"rohitsingh/misty-go/clients"

	"github.com/gorilla/mux"
	"github.com/rohitkochhar/reed-http-utills"
)

/*
- Entrypoint listens for PUT requests on an exposed IP and forwards
  the request to the misty broker that has services listening to different
  topics
- For simplicity, simple round-robin load balancing is used to route between
  three listening clients
*/

var packageCount int

func main() {
	// Listen on :2395/ for incoming PUT requests
	r := mux.NewRouter()
	r.HandleFunc("/", handlePut).Methods(http.MethodPut)
	log.Println("Creating entrypoint...")
	if err := http.ListenAndServe(
		":2395",
		r,
	); err != nil {
		os.Exit(1)
	}
}

func handlePut(w http.ResponseWriter, r *http.Request) {
	// The entrypoint publishes the information, but does
	// not request any information back, so we only need
	// to create a Publisher instance
	publisher := clients.NewPublisher()
	// Get the message from the body
	message, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		reed.ReplyError(w, r, http.StatusInternalServerError,
			"could not accept packet")
	}
	// Get the client that we are sending the information to
	targetClient := packageCount % 3
	packageCount++
	// Send a message to the logger to let it know we got
	// a message and we are now sending it to a specific client
	loggerMsg := fmt.Sprintf("Received message=\"%s\", forwarding to service-%d",
		message, targetClient,
	)
	// Send the message to the logger topic
	publisher.Publish("broker", 1315, "/logs", loggerMsg)
	// Send the message to the appropriate client
	publisher.Publish("broker", 1315, fmt.Sprintf("/service-%d", targetClient), string(message))
	// Reply back to the sender that the package was received
	// reed.ReplyError(w, r, http.StatusOK, "acknowledge")
}
