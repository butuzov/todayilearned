package httptesting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_WithMockEndpoint(t *testing.T) {
	// mock server for testing clients etc.
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`["foo", "bar"]`))
		}))
	defer server.Close()

	// requesting data with client server.URL
	client := New(server.URL)
	data, err := client.GetBooksEndpoint()
	assert.Equal(t, `["foo", "bar"]`, string(data))
	assert.NoError(t, err)
}
