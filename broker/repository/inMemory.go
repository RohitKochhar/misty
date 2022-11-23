package repository

import (
	"errors"
	"fmt"
	"sync"
)

var (
	ErrListenerExists         = errors.New("listener already exists in broker list")
	ErrListenerNotFound       = errors.New("listener not found in broker list")
	ErrTopicAlreadySubscribed = errors.New("listener is already subscribed to topic")
)

// inMemoryRepository type meets the Repository interface and only
// persists memory within local memory, meaning data is lost when
// server closes.
type inMemoryRepository struct {
	sync.RWMutex
	listeners map[string][]string // Maps listener (i.e. localhost:1111) to all topics ["/topic1", "/topic2", ... ]
}

// NewInMemoryRepository is the constructor for inMemoryRepository
// that just returns an empty inMemoryRepository object
func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		listeners: make(map[string][]string),
	}
}

// NewListener method adds a listener string (i.e. "localhost:1111")
// to the list of listeners, but does not attach any topics yet
func (r *inMemoryRepository) NewListener(listenerAddr string) error {
	r.Lock()
	defer r.Unlock()
	// Check that the listener isn't already in the repository
	if _, ok := r.listeners[listenerAddr]; ok {
		return fmt.Errorf("%w: %s", ErrListenerExists, listenerAddr)
	}
	// If it doesn't we can add it
	r.listeners[listenerAddr] = []string{}
	return nil
}

// AddTopicToListener method adds a topic to an existing listener,
// returns an error if the Listener is not found in the list
func (r *inMemoryRepository) AddTopicToListener(listenerAddr string, topic string) error {
	r.Lock()
	defer r.Unlock()
	// Check that the listener is in the repository
	if _, ok := r.listeners[listenerAddr]; !ok {
		return fmt.Errorf("%w: %s", ErrListenerNotFound, listenerAddr)
	}
	// Check if the topic is already in the listener's subscribed list
	for _, t := range r.listeners[listenerAddr] {
		if t == topic {
			return fmt.Errorf("%w: %s", ErrTopicAlreadySubscribed, topic)
		}
	}
	// If it exists, we can append the new topic to it's topic list
	r.listeners[listenerAddr] = append(r.listeners[listenerAddr], topic)
	return nil
}

// GetTopicListeners method returns all listener addresses that
// have subscribed to a given topic
func (r *inMemoryRepository) GetTopicListeners(topic string) ([]string, error) {
	// ToDo: Reduce complexity here (currently = #listeners * #topicsPerListener)
	// ToDo: Add some error handling here
	var topicListeners []string
	r.RLock()
	defer r.RUnlock()
	// Iterate through each listener in the repository
	for l, subscribedTopics := range r.listeners {
		// Check if the given topic is in the list
		for _, t := range subscribedTopics {
			if t == topic {
				topicListeners = append(topicListeners, l)
				break
			}
		}
	}
	return topicListeners, nil
}

// DeleteListener method removes a listener from the list
func (r *inMemoryRepository) DeleteListener(listenerAddr string) error {
	r.Lock()
	defer r.Unlock()
	delete(r.listeners, listenerAddr)
	return nil
}
