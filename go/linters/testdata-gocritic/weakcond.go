package testdata_gocritic

// Example for the rule: weakCond

func _(xs []*int) bool {
	return xs != nil && xs[0] != nil
}
