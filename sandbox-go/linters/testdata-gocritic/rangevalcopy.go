package testdata_gocritic

// Example for the rule: rangeValCopy

func _() {
	xs := make([][1024]byte, 512)
	for _, x := range xs {
		// Loop body.
		_ = x
	}
}
