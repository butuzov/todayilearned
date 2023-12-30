// file-header-check

package testdata

import (
	"log"
)

// Test IF and Boolean expr
func f(x int) bool { // MATCH /function f has cognitive complexity 3 (> max enabled 0)/
	if x > 0 && true || false { // +3
		return true
	} else {
		log.Printf("non-positive x: %d", x)
	}
	return false
}
