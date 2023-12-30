// file-header-check
package testdata

import "fmt"

func foo() {
	mySlice := []string{"A", "B", "C"}
	for index, value := range mySlice {
		go func() {
			fmt.Printf("Index: %d\n", index) // MATCH /loop variable index captured by func literal/
			fmt.Printf("Value: %s\n", value) // MATCH /loop variable value captured by func literal/
		}()
	}
}
