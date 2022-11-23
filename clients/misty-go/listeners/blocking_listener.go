package listeners

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
)

// Listener type implements a Listener interface that
// can only listen and reply to messages synchronously, blocking
// any additional client processes from running while the listener runs
type Listener struct {
	listenerUrl  string   // URL that the broker sends requests to
	listenerPort int      // Port to listen for incoming requests on
	brokerUrl    string   // URL to send broker requests to (no trailing "/")
	topics       []string // A list of all topics the client is subscribed to

	errCh  chan error     // Channel to receive errors on
	exitCh chan os.Signal // Channel to receive exit signals on
}

func NewListener(host string, port int) *Listener {
	// Define default parameters if empty vars are used
	if host == "" {
		host = "localhost"
	}
	if port == 0 {
		port = 1111
	}
	return &Listener{
		listenerUrl:  fmt.Sprintf("http://%s:%d", host, port),
		listenerPort: port,
	}
}

// Connect method sends a request to the misty broker to initiate a connection
func (l *Listener) Connect(host string, port int) error {
	// Send a PUT request to the broker to add this client to broker list
	httpUrl := fmt.Sprintf("http://%s:%d/listeners/_/add", host, port)
	if err := utils.PutString(httpUrl, l.listenerUrl, http.StatusAccepted); err != nil {
		return err
	}
	// Add the information about the broker to the listener object
	l.brokerUrl = fmt.Sprintf("http://%s:%d", host, port)
	return nil
}

// Close method sends a DELETE request to the misty broker to
// remove the client from the broker list
func (l *Listener) Close() error {
	// Send a DELETE request to the broker to delete this client from the broker list
	httpUrl := l.brokerUrl + "/listeners/delete"
	if err := utils.DeleteString(httpUrl, l.listenerUrl, http.StatusOK); err != nil {
		return err
	}
	return nil
}

// Subscribe method connects the listener to a topic
func (l *Listener) Subscribe(topic string) error {
	// Sanitize the topic to make sure it will fit the HTTP protocol
	sanitizedTopic, err := utils.SanitizeTopic(topic)
	if err != nil {
		return fmt.Errorf("error while sanitizing topic: %q", err)
	}
	// Send a PUT request to the broker to add topic to this clients subscription list
	httpUrl := l.brokerUrl + "/listeners" + sanitizedTopic + "/subscribe"
	if err := utils.PutString(httpUrl, l.listenerUrl, http.StatusAccepted); err != nil {
		return fmt.Errorf("error while trying to subscribe to topic=%s: %q", topic, err)
	}
	l.topics = append(l.topics, sanitizedTopic)
	return nil
}

// Listen method listens for published messages on all the subscribed topics
func (l *Listener) Listen() error {
	// Start an asynchronous HTTP server to receive messages from the broker
	r := mux.NewRouter()
	// Add a handler for each topic the listener is subscribed to
	for _, topic := range l.topics {
		r.HandleFunc(topic, onMessageReceive)
	}
	// Create goroutine resources
	l.errCh = make(chan error)
	l.exitCh = make(chan os.Signal, 1)
	signal.Notify(l.exitCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", l.listenerPort), r); err != nil {
			l.errCh <- err
		}
	}()
	select {
	case err := <-l.errCh:
		log.Println(err)
		if err := l.Close(); err != nil {
			return fmt.Errorf("couldn't remove listener from broker list due to error: %q", err)
		}
		return err
	case <-l.exitCh:
		log.Println("Received termination signal, closing listener...")
		if err := l.Close(); err != nil {
			return fmt.Errorf("couldn't remove listener from broker list due to error: %q", err)
		}
		log.Println("Successfully removed listener from broker list")
	}
	return nil
}

// ToDo: Remove this and make it a user defined function
func onMessageReceive(w http.ResponseWriter, r *http.Request) {
	// Get the value from the request
	message, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.ReplyError(w, r, http.StatusInternalServerError, "could not accept packet")
	}
	utils.ReplyTextContent(w, r, http.StatusAccepted, "accepted packet")
	log.Printf("[RECEIVED ON %s]: %s", r.URL.Path, string(message))
}
