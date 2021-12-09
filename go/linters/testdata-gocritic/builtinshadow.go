package testdata_gocritic

// Example for the rule: builtinShadow

func _() int {
	len := 10
	return len
}
