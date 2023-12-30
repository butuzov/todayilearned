package testdata_gocritic

// Example for the rule: underef

type ts struct {
	field int
}

const five int = 5

func (k *ts) _(a *[]int) {
	(*k).field = five
	_ = (*a)[5] // only if a is array
}
