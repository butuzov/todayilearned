package testdata_gocritic

// Example for the rule: typeAssertChain

type (
	T1 int
	T2 int
	T3 int
)

func _(v interface{}) {

	if x, ok := v.(T1); ok {
		// Code A, uses x.
		_ = x
	} else if x, ok := v.(T2); ok {
		// Code B, uses x.
		_ = x
	} else if x, ok := v.(T3); ok {
		// Code C, uses x.
		_ = x
	}

}
