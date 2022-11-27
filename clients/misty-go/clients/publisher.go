package clients

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rohitkochhar/reed-http-utills"
)

// Publisher type implements a Publisher interface
// that can connect and send messages to a broker
type Publisher struct {
}

func NewPublisher() *Publisher {
	return &Publisher{}
}

// Publish sends a PUT request to the server containing some data
func (p *Publisher) Publish(host string, port int, topic string, message string) error {
	// Sanitize the topic
	topic, err := SanitizeTopic(topic)
	if err != nil {
		return err
	}
	// Send a PUT request to the specified broker
	httpUrl := fmt.Sprintf("http://%s:%d/topic%s", host, port, topic)
	if err := reed.PutString(httpUrl, message, []int{http.StatusOK}); err != nil {
		log.Printf("error while trying to publish %s on %s", message, topic)
		return err
	}
	log.Printf("Successfully published %s on %s", message, topic)
	return nil
}
