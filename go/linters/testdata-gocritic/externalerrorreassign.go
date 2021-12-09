package testdata_gocritic

import "io"

func isEOF(e error) bool {
	io.EOF = nil
	return e == io.EOF
}
