package testdata

import "fmt"

// PROJECT URL: https://github.com/mvdan/gofumpt
func GofumptNewLine() {
	fmt.Println( "foo" ) // ERROR "File is not `gofumpt`-ed"
}
