package httptesting

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gavv/httpexpect"
	"github.com/stretchr/testify/assert"
)

// running midleware (regular http.handler) viv

// assertify...
func TestMiddlewareWithAssert(t *testing.T) {
	handler := StopWatch(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})).(http.HandlerFunc)

	assert.HTTPBodyContains(t, handler, "GET", "/", nil, "ok")
	assert.HTTPSuccess(t, handler, "GET", "/", nil, "ok")
}

// regular recorder...
func TestMiddlewareWithRecorder(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	// running request
	StopWatch(handler).ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "ok")
}

// httpexpect
func TestMiddlewareWithHTTPEXpect(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	// run server using httptest
	server := httptest.NewServer(StopWatch(handler))
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)
	e.GET("/").Expect().Status(http.StatusOK).Body().Equal("ok")
}
