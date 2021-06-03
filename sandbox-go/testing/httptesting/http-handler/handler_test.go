package httphandler

import (
	"httptesting"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestHttpHandler(t *testing.T) {
	suite.Run(t, &httptesting.HandlerTestSuite{
		Handler: Handler{
			Users: map[string]string{
				"foo": "bar",
			},
		},
	})
}
