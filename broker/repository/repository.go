package repository

type Repository interface {
	// NewListener method adds a listener string (i.e. "localhost:1111")
	// to the list of listeners, but does not attach any topics yet
	NewListener(listenerAddr string) error
	// AddTopicToListener method adds a topic to an existing listener,
	// returns an error if the Listener is not found in the list
	AddTopicToListener(listenerAddr string, topic string) error
	// GetTopicListeners method returns all listener addresses that
	// have subscribed to a given topic
	GetTopicListeners(topic string) ([]string, error)
	// DeleteListener method removes a listener from the list
	DeleteListener(listenerAddr string) error
	// GetAllListeners method returns all listeners in the repository
	GetAllListeners() ([]string, error)
}
