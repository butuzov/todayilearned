package testdata

import (
	"github.com/pkg/errors"
)

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
