package testdata_gocritic

// Example for the rule: sloppyLen

func _(arr []int) bool {
	b1 := len(arr) >= 0 // Sloppy
	b2 := len(arr) <= 0 // Sloppy
	b3 := len(arr) < 0  // Doesn't make sense at all
	return b1 && b2 && b3
}
