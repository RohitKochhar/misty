package service_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"rohitsingh/misty-broker/service"
	"strings"
	"testing"
)

// setupAPI is a helper function that sets up
// the API for the tests, providing a cleanup function too
func setupAPI(t *testing.T) (string, func()) {
	t.Helper() // Mark the function as test helper
	ts := httptest.NewServer(service.NewMux())
	return ts.URL, func() {
		ts.Close()
	}
}

// getHelper wraps the Get function in additional logic to
// assist with testing ease and clarity
func getHelper(t *testing.T, getUrl string, expBody string, expCode int) (r *http.Response) {
	r, err := http.Get(getUrl)
	if err != nil {
		t.Fatalf("error while sending GET request: %q", err)
	}
	// Check if the return code is what we expected
	if r.StatusCode != expCode {
		t.Fatalf("Expected %q, got %q.", http.StatusText(expCode),
			http.StatusText(r.StatusCode))
	}
	defer r.Body.Close()
	// We might not be expecting content
	if expBody != "" || expCode == http.StatusNotFound {
		// The result of GET from the server should always be in
		// plain text
		if !strings.Contains(r.Header.Get("content-Type"), "text/plain") {
			t.Fatalf("unsupported Content-Type: %q", r.Header.Get("Content-Type"))
		}
		// Check that we have the content that we expected
		var body []byte
		if body, err = io.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(string(body), expBody) {
			t.Fatalf("Expected %q, got %q.", expBody, string(body))
		}
	}

	return r
}

// TestConnectionchecks if we can ping the server
// and get a response back
func TestConnection(t *testing.T) {
	// Initialize a test server
	httpUrl, cleanup := setupAPI(t)
	defer cleanup()
	// Send a GET request to the root of the
	// misty server
	_ = getHelper(t, httpUrl, "misty server ready to accept connections at", http.StatusOK)
}
