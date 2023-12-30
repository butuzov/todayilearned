package testdata_gocritic

// Example for the rule: emptyStringTest

func foobar10(s string) bool {
	return len(s) == 0
}
