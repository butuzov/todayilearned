package testdata_gocritic

// Example for the rule: commentedOutImport

import (
	"fmt"
	// "os"
)

func _(n int) {
	fmt.Sprintf("%d", n)
}
