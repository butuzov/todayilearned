package testdata

import (
	"os"

	"github.com/pkg/errors"
)

// OK_with_config
// PROJECT URL:
// https://github.com/gostaticanalysis/nilerr

func j() (interface{}, error) {
	do := func() error { return errors.New("damn") }

	if err := do(); err != nil {
		return nil, nil // want "error is not nil (line 23) but it returns nil"
	}

	if err := do(); err != nil {
		return nil, err
	}

	if err := do(); err != nil {
		return err, nil // want "error is not nil (line 31) but it returns nil"
	}

	if err := do(); err != nil {
		return err, err
	}

	return nil, nil
}

// works?
func nilErr2() error {
	err := nilErrDo()
	if err == nil {
		return err // want `error is nil \(line 16\) but it returns error`
	}

	return nil
}

func nilErrDo() error {
	return os.ErrNotExist
}
