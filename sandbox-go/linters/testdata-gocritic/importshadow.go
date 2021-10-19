package testdata_gocritic

import "path/filepath"

// Example for the rule: importShadow

func getBadPettern() error {
	return filepath.ErrBadPattern
}

func _() {
	// "path/filepath" is imported.
	filepath := "foo.txt"
	_ = filepath
}
