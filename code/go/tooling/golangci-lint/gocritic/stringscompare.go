package testdata_gocritic

import "strings"

func _(a, b string) bool {
	return strings.Compare(a, b) == 0
}
