package testdata_gocritic

// Example for the rule: offBy1

func _(n int) int {
	var xs = make([]int, n)
	return xs[len(xs)]
}
