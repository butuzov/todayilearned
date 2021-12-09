package testdata_gocritic

// Example for the rule: badCond

func _(n int) {
	for i := 0; i > n; i++ {
		// do nothing
	}
}
