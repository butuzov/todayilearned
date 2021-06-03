package testdata_gocritic

// Example for the rule: yodaStyleExpr

type T9 struct {
}

func _(ptr *T9) bool {
	return nil != ptr
}
