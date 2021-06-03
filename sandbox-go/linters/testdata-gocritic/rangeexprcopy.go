package testdata_gocritic

// Example for the rule: rangeExprCopy

func _(n int) {
	var xs [2048]byte
	for _, x := range xs { // Copies 2048 bytes
		// Loop body.
		_ = x
	}
}
