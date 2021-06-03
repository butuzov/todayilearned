package testdata_gocritic

import "io"

func _(r io.Reader) interface{} {
	return r.(interface{})
}
