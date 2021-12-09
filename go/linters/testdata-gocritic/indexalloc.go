package testdata_gocritic

import "strings"

// Example for the rule: indexAlloc

func _(x []byte, y string) {
	strings.Index(string(x), y)
}
