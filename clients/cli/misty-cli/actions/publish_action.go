package actions

import (
	"bytes"
	"fmt"
	"net/http"
)

// Publish sends a PUT request to send a given
// message to a given topic and returns the
// response error status
func Publish(host string, port int, topic string, message string) error {
	// PUT the message on the corresponding topic
	// ToDo: Add topic sanitizing / function to create HTTP URL with rules
	httpUrl := fmt.Sprintf("http://%s:%d%s", host, port, topic)
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
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error while making PUT request: %s", http.StatusText(resp.StatusCode))
	}
	fmt.Printf("[PUBLISH] %s --> %s\n", message, httpUrl)
	return nil
}
