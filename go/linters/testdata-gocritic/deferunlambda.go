package testdata_gocritic

func fExample10() {}

func _() {
	defer func() { fExample10() }()
}
