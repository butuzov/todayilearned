package httptesting

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// using assert to test simple handlerfunc

func TestAssertHTTPBodyContains(t *testing.T) {
	// mock server for testing clients etc.
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	assert.HTTPBodyContains(t, handler, "GET", "/", nil, "ok")
}
