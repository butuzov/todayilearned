package testdata_gocritic

// Example for the rule: switchTrue

func _(x, y int) {
	switch true {
	case x > y:
	}
}
