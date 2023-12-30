package testdata_gocritic

// Example for thehe rule: appendAssign

func _(n int) bool {
	var (
		pos, neg []int
	)
	pos = append(neg, n)
	neg = append(pos, n)

	return len(pos) != len(neg)
}
