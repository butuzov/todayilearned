package testdata_gocritic

import (
	"io"
	"io/ioutil"
)

func readall(r io.Reader) []byte {
	b, _ := ioutil.ReadAll(r)
	return b
}
