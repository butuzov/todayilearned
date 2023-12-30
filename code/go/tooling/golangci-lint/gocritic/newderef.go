package testdata_gocritic

// Example for the rule: newDeref

func _() bool {
	x := *new(bool)
	return x
}
