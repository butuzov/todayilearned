package testdata_gocritic

// Example for the rule: hexLiteral

func _() int {
	x := 0x12
	y := 0xfF
	return x + y
}
