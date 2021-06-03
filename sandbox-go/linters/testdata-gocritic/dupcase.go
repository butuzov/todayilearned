package testdata_gocritic

// Example for the rule: dupCase

func _(n int) {
	var ys = make([]bool, n)
	switch n > 10 {
	case ys[0], ys[1], ys[2], ys[0], ys[4]:

	}
}
