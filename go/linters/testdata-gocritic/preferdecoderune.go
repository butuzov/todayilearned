package testdata_gocritic

// Example for the rule: preferDecodeRune
var (
	s = "f"
	r = []rune(s)[0]
)
