package testdata_gocritic

// Example for the rule: dupImport

import (
	"fmt"
	priting "fmt" // Imported the second time
)

func _(n int) bool {
	return priting.Sprintf("%d", n) == fmt.Sprintf("%d", n)
}
