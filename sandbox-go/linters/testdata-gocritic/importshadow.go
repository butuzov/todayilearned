package testdata_gocritic

// Example for the rule: importShadow

func _() {
	// "path/filepath" is imported.
	filepath := "foo.txt"
	_ = filepath
}
