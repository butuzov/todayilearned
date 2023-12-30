package testdata

import "encoding/json"

// PROJECT URL: https://github.com/tomarrell/wrapcheck
func _() error {
	_, err := json.Marshal(struct{}{})
	if err != nil {
		return err // ERROR "error returned from external package is unwrapped"
	}

	return nil
}
