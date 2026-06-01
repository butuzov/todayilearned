<!-- weight: 500 -->
<!-- menu: <code>httptesting</code>  -->
<!-- github: https://github.com/butuzov/todayilearned/tree/main/coding/go/testing/httptesting -->

# `httptesting`

## `testify/assert`

```go
package hello_httptesting

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// using assert to test simple handlerfunc

func TestAssertHTTPBodyContains_HelloWorld(t *testing.T) {
	// mock server for testing clients etc.
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HelloWorld"))
	})

	assert.HTTPBodyContains(t, handler, "GET", "/", nil, "HelloWorld")
}
```

## `gavv/httpexpect`

```go
import "github.com/gavv/httpexpect"

func TestMiddlewareWithHTTPEXpect(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HelloWorld"))
	})

	// run server using httptest
	server := httptest.NewServer(StopWatch(handler))
	defer server.Close()

	// create httpexpect instance
	e := httpexpect.New(t, server.URL)
	e.GET("/").Expect().Status(http.StatusOK).Body().Equal("HelloWorld")
}
```

## `httptest.NewRecorder`

```go
import "net/http/httptest"

func TestMiddlewareWithRecorder(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("HelloWorld"))
	})

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	// running request
	StopWatch(handler).ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.String(), "HelloWorld")
}

```

## `HandlerTestSuite`

{{% list "middleware_stopwatch.go,handler_testing_suite.go,middleware_panic.go" %}}

```

```
