package testdata_gocritic

// Example for the rule: initClause

func _(n int) {
	if sideEffect(); n > 10 {
	}
}

func sideEffect() {}
