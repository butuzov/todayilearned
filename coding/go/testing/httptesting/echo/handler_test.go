package testecho

import (
	"httptesting"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestEchoHandler(t *testing.T) {
	suite.Run(t, &httptesting.HandlerTestSuite{
		Handler: Handler(map[string]string{"foo": "bar"}),
	})
}
