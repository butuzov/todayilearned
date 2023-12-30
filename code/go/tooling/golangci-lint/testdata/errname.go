package testdata

import (
	"fmt"
)

var (
	EOF          = fmt.Errorf("end of file")
	ErrEndOfFile = fmt.Errorf("end of file")
	errEndOfFile = fmt.Errorf("end of file")

	EndOfFileError = fmt.Errorf("end of file") // want "the variable name `EndOfFileError` should conform to the `ErrXxx` format"
	ErrorEndOfFile = fmt.Errorf("end of file") // want "the variable name `ErrorEndOfFile` should conform to the `ErrXxx` format"
	EndOfFileErr   = fmt.Errorf("end of file") // want "the variable name `EndOfFileErr` should conform to the `ErrXxx` format"
	endOfFileError = fmt.Errorf("end of file") // want "the variable name `endOfFileError` should conform to the `errXxx` format"
	errorEndOfFile = fmt.Errorf("end of file") // want "the variable name `errorEndOfFile` should conform to the `errXxx` format"
)
