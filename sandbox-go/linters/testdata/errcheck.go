package testdata

import "github.com/pkg/errors"

// PROJECT URL:  https://github.com/gordonklaus/ineffassign
func _() {
	returnError()
}

func returnError() error {
	return errors.New("sss")
}
