package testdata_gocritic

// Example for the rule: elseif

func _(n int) {
	if n > 10 {
	} else {
		if x := n > 10; x {
		}
	}

}
