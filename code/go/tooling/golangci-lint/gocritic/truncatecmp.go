package testdata_gocritic

// Example for the rule: truncateCmp

func _(x int32, y int16) bool {
	return int16(x) < y
}
