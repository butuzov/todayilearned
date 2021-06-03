package httptesting

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_WithMockEndpoint(t *testing.T) {
	// mock server for testing clients etc.
	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			f, _ := os.Open("testdata/foobar.json")
			defer f.Close()

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			io.Copy(w, f)
		}))
	defer server.Close()

	// requesting data with client server.URL
	client := New(server.URL)
	data, err := client.GetBooksEndpoint()
	assert.Equal(t, `["foo", "bar"]`, string(data))
	assert.NoError(t, err)
}
