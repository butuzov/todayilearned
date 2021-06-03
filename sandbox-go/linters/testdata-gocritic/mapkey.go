package testdata_gocritic

// Example for the rule: mapKey

func _() {
	_ = map[string]int{
		" foo": 1,
		"bar ": 2,
	}
}
