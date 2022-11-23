package service

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rohitsingh/misty-broker/repository"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var repo repository.Repository

// Run configures a HTTP server and listens for incoming requests
func Run() error {
	// Create a repository to contain listener data
	// ToDo: Add handling to change repo type depending on config
	repo = repository.NewInMemoryRepository()
	// Load configuration variables from viper
	host := viper.GetString("host")
	port := viper.GetInt("port")
	log.Printf("Creating misty broker at http://%s:%d...\n", host, port)
	// Create a mux and listen and serve on it
	r := NewMux()
	// ToDo: Add a configuration parameter that specifies HTTPS/HTTP
	// Create goroutine resources
	errCh := make(chan error)
	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt, syscall.SIGTERM)
	// Run server as a goroutine
	go func() {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), r); err != nil {
			errCh <- err
		}
	}()
	select {
	case err := <-errCh:
		log.Println(err)
		if err := CloseConnections(); err != nil {
			return fmt.Errorf("couldn't send termination signal to listeners due to error: %q", err)
		}
		return err
	case <-exitCh:
		log.Println("Received termination signal, closing listener...")
		if err := CloseConnections(); err != nil {
			return fmt.Errorf("couldn't send termination signal to listeners due to error: %q", err)
		}
		log.Println("Successfully terminated all connections, shutting down now")
		os.Exit(1)
	}
	return nil
}

// NewMux creates a HTTP router and attaches handlers to it
func NewMux() http.Handler {
	r := mux.NewRouter()
	// Root path is used as a liveness check
	r.HandleFunc("/", rootHandler).Methods(http.MethodGet)
	r.HandleFunc("/topic/{topic}", publishHandler).Methods(http.MethodPut)
	r.HandleFunc("/listeners/{topic}/add", addListenerHandler).Methods(http.MethodPut)
	r.HandleFunc("/listeners/delete", deleteListenerHandler).Methods(http.MethodDelete)
	r.HandleFunc("/listeners/{topic}/subscribe", listenerSubscribeHandler).Methods(http.MethodPut)
	return r
}

// replyTextContent wraps text content in a HTTP response and sends it
func replyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(content + "\n"))
}

// replyError wraps text content in an HTTP error response and sends it
func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error: %d %s", r.URL, r.Method, status, message)
	http.Error(w, http.StatusText(status), status)
}
