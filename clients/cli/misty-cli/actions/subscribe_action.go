package actions

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	utils "rohitsingh/misty-utils"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

// SubscribeAction creates a server to accept messages by the broker
func SubscribeAction(brokerHost string, brokerPort int, topic string) error {
	// Sanitize the topic
	topic, err := utils.SanitizeTopic(topic)
	if err != nil {
		return err
	}
	// Define the parameters used by the listener as a HTTP server
	var (
		listenerHost string
		listenerPort int
	)
	if listenerHost = viper.GetString("host"); listenerHost == "" {
		listenerHost = "localhost"
	}
	if listenerPort = viper.GetInt("port"); listenerPort == 0 {
		listenerPort = 1111
	}
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
	r.HandleFunc(topic, subscribeHandler).Methods(http.MethodPut)
	log.Printf("listening for messages on %s...\n", topic)
	// Create goroutine resources
	errCh := make(chan error)
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", listenerPort), r); err != nil {
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		if err := removeListener(brokerHost, brokerPort, listenerHost, listenerPort); err != nil {
			log.Panicf("couldn't remove listener from broker list due to error: %q", err)
		}
		log.Println(err)
	case <-exitCh:
		log.Println("Received termination signal, closing listener...")
		if err := removeListener(brokerHost, brokerPort, listenerHost, listenerPort); err != nil {
			log.Panicf("couldn't remove listener from broker list due to error: %q", err)
		}
		log.Println("Removed listener from broker list")
	}
	return nil
}

// subscribeHandler handles incoming PUT requests to a topic
func subscribeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the value from the request
	message, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.ReplyError(w, r, http.StatusInternalServerError, "could not accept packet")
	}
	utils.ReplyTextContent(w, r, http.StatusAccepted, "accepted packet")
	log.Printf("[RECEIVED ON %s]: %s", r.URL.Path, string(message))
}

// requestSubscribe sends a message to the broker letting it know
// that this client wants to receive messages for a given topic
func requestSubscribe(brokerHost string, brokerPort int, listenerHost string, listenerPort int, topic string) error {
	// Sanitize the topic
	topic, err := utils.SanitizeTopic(topic)
	if err != nil {
		return err
	}
	// Send a PUT request to the broker containing listener server information
	httpUrl := fmt.Sprintf("http://%s:%d/listeners%s/add", brokerHost, brokerPort, topic)
	message := fmt.Sprintf("http://%s:%d", listenerHost, listenerPort)
	if err := utils.PutString(httpUrl, message, http.StatusAccepted); err != nil {
		return err
	}
	return nil
}

// removeListener sends a request to the server to remove the listener
// from its data store
func removeListener(brokerHost string, brokerPort int, listenerHost string, listenerPort int) error {
	// Send a DELETE request to the server
	httpUrl := fmt.Sprintf("http://%s:%d/listeners/delete", brokerHost, brokerPort)
	message := fmt.Sprintf("%s:%d", listenerHost, listenerPort)
	if err := utils.DeleteString(httpUrl, message, http.StatusOK); err != nil {
		return err
	}
	return nil
}
