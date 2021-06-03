package testdata_gocritic

// Example for the rule: sloppyReassign

func _(n int) error {
	var err error
	if err = func() error { return nil }(); err != nil {
		return err
	}
	return nil
}
