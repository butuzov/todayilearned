package testdata_gocritic

func expr() error {
	return nil
}

func _(err2 error) {
	if err := expr(); err2 != nil {
		_ = err
	}
}
