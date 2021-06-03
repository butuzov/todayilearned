package testdata_gocritic

// Example for the rule: valSwap

func _(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}
