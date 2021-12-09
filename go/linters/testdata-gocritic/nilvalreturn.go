package testdata_gocritic

import "errors"

// Example for the rule: nilValReturn

func _() error {

	err := errors.New("oops")

	if err == nil {
		return err
	}

	return nil
}
