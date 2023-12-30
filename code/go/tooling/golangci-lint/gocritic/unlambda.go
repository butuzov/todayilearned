package testdata_gocritic

// Example for the rule: unlambda

func _(n int) int {
	return func(x int) int { return f8(x) }(n)
}

func f8(x int) int {
	return x * x
}
