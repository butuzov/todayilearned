package testdata_gocritic

import (
	"log"
	"os"
)

// Example for the rule: exitAfterDefer

func _(name string) {

	defer os.Remove(name)
	if _, err := os.Stat(name); os.IsNotExist(err) {
		log.Fatalf("something bad happened")
	}
}
