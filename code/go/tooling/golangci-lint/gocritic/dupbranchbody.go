package testdata_gocritic

// Example for the rule: dupBranchBody

func _(n int) {
	if n > 0 {
		println("cond=true")
	} else {
		println("cond=true")
	}
}
