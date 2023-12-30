package testecho

import (
	"context"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHTTPErrorHandler(t *testing.T) {
	e := echo.New()
	defer func() { _ = e.Close() }()

	tests := map[string]struct {
		err     error
		status  int
		message string
	}{
		"context.DeadlineExceeded": {
			context.DeadlineExceeded,
			408,
			`{"message":"context deadline exceeded"}` + "\n",
		},
		"context.Canceled": {
			context.Canceled,
			408,
			`{"message":"context canceled"}` + "\n",
		},
		`echo.HTTPError`: {
			&echo.HTTPError{
				Code:    418,
				Message: "I'm a teapot",
			},
			418,
			`{"message":"I'm a teapot"}` + "\n",
		},

		`regular_error`: {
			errors.New("I'm a teapot"),
			http.StatusInternalServerError,
			`{"message":"I'm a teapot"}` + "\n",
		},
	}

	for name, test := range tests {
		name, test := name, test

		t.Run(name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/", nil)
			rec := httptest.NewRecorder()

			errorHandler(test.err, e.NewContext(req, rec))

			b, _ := ioutil.ReadAll(rec.Body)

			assert.Equal(t, test.status, rec.Code, "Error Status Code not match")
			assert.Equal(t, test.message, string(b), "Error message not match")
		})
	}
}

type echoError struct {
	Message string `json:"message"`
}

func (e echoError) Error() string {
	return e.Message
}

func errorHandler(err error, c echo.Context) {
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		c.JSON(http.StatusRequestTimeout, echoError{err.Error()})
		return
	}

	var echoErr *echo.HTTPError
	if errors.As(err, &echoErr) {
		c.JSON(echoErr.Code, echoErr)
		return
	}

	// Add another Error handling here
	c.JSON(http.StatusInternalServerError, echoError{err.Error()})
}
