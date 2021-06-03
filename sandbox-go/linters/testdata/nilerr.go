package testdata

import "github.com/pkg/errors"

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
