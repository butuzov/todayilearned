package testdata_gocritic

// Example for the rule: dupSubExpr

var s1, s2 []int

func _(n int) bool {
	return s1[n] > s1[n]
}
