package testdata_gocritic

import "path/filepath"

func _() {
	_ = filepath.Join("dir/", "test.txt")
}
