package testdata_gocritic

import "strings"

// Example for the rule: badCall

func _(s, from, to string) string {
	return strings.Replace(s, from, to, 0)
}
