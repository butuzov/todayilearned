package testdata

import "time"

func _(t *time.Duration) error {
	if t == nil {
		return nil
	} else { // ERROR "indent-error-flow: if block ends with a return statement, .*"
		return nil
	}
}
