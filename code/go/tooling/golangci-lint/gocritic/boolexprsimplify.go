package testdata_gocritic

// Example for the rule: boolExprSimplify

func _(max, min int, y, x bool) bool {
	a := !(max >= min)
	b := !(x) == !(y)

	return a && b
}
