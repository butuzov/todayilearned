package testdata_gocritic

// Example for the rule: singleCaseSwitch

func _(x interface{}) {
	switch x.(type) {
	case int:
		func() {}()
	}
}
