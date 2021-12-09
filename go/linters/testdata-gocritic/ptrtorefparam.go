package testdata_gocritic

// Example for the rule: ptrToRefParam

func _(m *map[string]int) *chan *int {
	return nil
}
