package testdata_gocritic

// Example for the rule: unlabelStmt

func _(xs []int) {
derp:
	for x := range xs {
		if x == 0 {
			break derp
		}
	}
}
