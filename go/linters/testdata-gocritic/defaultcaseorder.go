package testdata_gocritic

// Example for the rule: defaultCaseOrder

func _(x, y int) {
	switch {
	case x > y:
		// ...
	default: // <- not the best position
		// ...
	case x == 10:
		// ...
	}
}
