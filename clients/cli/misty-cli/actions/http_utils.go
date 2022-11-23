package actions

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

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

// PutString wraps a HTTP Get request in a function for easier usage and clarity
func PutString(httpUrl string, message string, expCode int) error {
	req, err := http.NewRequest(
		http.MethodPut,
		httpUrl,
		bytes.NewBuffer([]byte(message)),
	)
	if err != nil {
		return fmt.Errorf("error while creating PUT request: %q", err)
	}
	req.Header.Set("Content-Type", "text/plain")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error while sending PUT request: %q", err)
	}
	if resp.StatusCode != expCode {
		return fmt.Errorf("error while making PUT request: %s", http.StatusText(resp.StatusCode))
	}
	return nil
}
