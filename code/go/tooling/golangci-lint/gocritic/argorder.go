package testdata_gocritic

import "strings"

// Example for the rule: argOrder

func _(s string) bool {
	return strings.HasPrefix("http://", s)
}
