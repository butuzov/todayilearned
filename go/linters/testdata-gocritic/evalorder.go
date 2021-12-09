package testdata_gocritic

// Example for the rule: evalOrder

func f1(n *int) error {
	return nil
}

func _(n int) (int, error) {
	return n, f1(&n)
}
