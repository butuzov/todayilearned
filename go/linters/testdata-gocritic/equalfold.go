package testdata_gocritic

import "strings"

// Example for the rule: equalFold

func _(x, y string) bool {
	return strings.ToLower(x) == strings.ToLower(y)
}
