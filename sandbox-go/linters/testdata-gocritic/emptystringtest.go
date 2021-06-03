package testdata_gocritic

// Example for the rule: emptyStringTest

func _(s string) bool {
	return len(s) == 0
}
