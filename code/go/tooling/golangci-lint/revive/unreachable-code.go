// file-header-check
package testdata

import (
	"fmt"
	"log"
)

func fooUnreachable() int {
	log.Fatalf("%s", "About to fail") // ignore
	return 0                          // MATCH /unreachable code after this statement/

	fmt.Println("unreachable")
	return 1
}
