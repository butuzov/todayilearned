package testdata_gocritic

// Example for the rule: dupArg

func _() bool {

	var dst, src []int

	copy(dst, dst)

	return len(src) == len(dst)

}
