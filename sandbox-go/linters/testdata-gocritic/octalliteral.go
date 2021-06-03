package testdata_gocritic

// Example for the rule: octalLiteral

func _() {
	f2(02)
}

func f2(n int) {}
