package testdata_gocritic

import "os"

func Remove(name string) {
	defer os.Remove(name)
	print("remove")
}
