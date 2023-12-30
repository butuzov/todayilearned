// file-header-check
package testdata

import (
	"log"
	"os"
)

func foo0() {
	os.Exit(1) // MATCH /calls to os.Exit only in main() or init() functions/
}

func init() {
	log.Fatal("v ...interface{}")
}
