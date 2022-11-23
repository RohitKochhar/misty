package repository_test

import (
	"errors"
	"rohitsingh/misty-broker/repository"
	"testing"
)

// TestInMemoryRepoIntegration performs an integration flow test on the
// inMemoryRepository object to ensure it works as intended
func TestInMemoryRepoIntegration(t *testing.T) {
	// Create a new inMemoryRepository
	repo := repository.NewInMemoryRepository()
	// Add a listener to the repository
	if err := repo.NewListener("localhost:1111"); err != nil {
		t.Fatalf("error while adding new listener to repository: %q", err)
	}
	// Add the same listener to the repository
	// expect to fail since the listener already exists
	if err := repo.NewListener("localhost:1111"); !errors.Is(err, repository.ErrListenerExists) {
		t.Fatalf("expected %q, instead got %q", repository.ErrListenerExists, err)
	}
	// Add second listener
	if err := repo.NewListener("localhost:1112"); err != nil {
		t.Fatalf("error while adding new listener to repository: %q", err)
	}
	// Add topic to both listeners
	if err := repo.AddTopicToListener("localhost:1111", "/topic"); err != nil {
		t.Fatalf("error while adding topic `%s` to %s: %q", "/topic", "localhost:1111", err)
	}
	if err := repo.AddTopicToListener("localhost:1112", "/topic"); err != nil {
		t.Fatalf("error while adding topic `%s` to %s: %q", "/topic", "localhost:1112", err)
	}
	// Check that the topic can't be added again
	if err := repo.AddTopicToListener("localhost:1111", "/topic"); !errors.Is(err, repository.ErrTopicAlreadySubscribed) {
		t.Fatalf("expected %q, instead got %q", repository.ErrTopicAlreadySubscribed, err)
	}
	if err := repo.AddTopicToListener("localhost:1112", "/topic"); !errors.Is(err, repository.ErrTopicAlreadySubscribed) {
		t.Fatalf("expected %q, instead got %q", repository.ErrTopicAlreadySubscribed, err)
	}
	// Check that we can get the listeners by querying the topic
	topicListeners, err := repo.GetTopicListeners("/topic")
	if err != nil {
		t.Fatalf("error while getting topic listeners for %s", "/topic")
	}
	if len(topicListeners) != 2 {
		t.Fatalf("expected topicListeners to contain %d listeners, instead got %d", 2, len(topicListeners))
	}
	// Delete the first listener
	if err := repo.DeleteListener("localhost:1111"); err != nil {
		t.Fatalf("error while trying to delete listener: %q", err)
	}
	// Check that we have only one listener on the topic now
	topicListeners, err = repo.GetTopicListeners("/topic")
	if err != nil {
		t.Fatalf("error while getting topic listeners for %s", "/topic")
	}
	if len(topicListeners) != 1 {
		t.Fatalf("expected topicListeners to contain %d listeners, instead got %d", 1, len(topicListeners))
	}
}
