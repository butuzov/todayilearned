package testdata_gocritic

// Example for the rule: nestingReduce

func loopWithIf(a []int) {
	for _, v := range a {
		/*! invert if cond, replace body with `continue`, move old body after the statement */
		if v == 5 {
			_ = v
			_ = v
			_ = v
			_ = v
			_ = v
			_ = v
		}
	}
}
