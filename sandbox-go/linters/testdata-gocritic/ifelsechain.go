package testdata_gocritic

// Example for the rule: ifElseChain

func _(n int) int {
	if n > 1 {
		n *= 10
	} else if n > 10 {
		n *= 100

	} else if n > 1100 {
		n *= 110
	}
	return n
}
