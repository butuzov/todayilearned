package testdata

import (
	"io/ioutil"
	"os"
)

// PROJECT URL:  https://github.com/mvdan/interfacer
func ProcessInput(f *os.File) int {
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return -1
	}
	return len(b)
}
