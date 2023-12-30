package testdata

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
)

// OK_with_config

// PROJECT URL:  https://github.com/polyfloyd/go-errorlint

func errorlint() error {
	errDefault := errors.New("default value")

	funcTestErrorlint := func() error {
		return errDefault
	}

	if err := funcTestErrorlint(); err == errDefault {
		return errors.Errorf("oh noes: %w", err)
	} else if err != nil {
		return errors.Errorf("not nilll: %v", err)
	}

	return nil
}

// example 2

var errLintFoo = errors.New("foo")

type errLintBar struct{}

func (*errLintBar) Error() string {
	return "bar"
}

func errorLintAll() {
	err := func() error { return nil }()
	if err == errLintFoo { // want "comparing with == will fail on wrapped errors. Use errors.Is to check for a specific error"
		log.Println("errCompare")
	}

	err = errors.New("oops")
	fmt.Errorf("error: %v", err) // want "non-wrapping format verb for fmt.Errorf. Use `%w` to format errors"

	switch err.(type) { // want "type switch on error will fail on wrapped errors. Use errors.As to check for specific errors"
	case *errLintBar:
		log.Println("errLintBar")
	}
}
