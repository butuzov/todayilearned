package testdata

import "fmt"

// PROJECT URL:
// https://github.com/gostaticanalysis/forcetypeassert

func forcetypeassertInvalid() {
	var a interface{}
	_ = a.(int) // ERROR "type assertion must be checked"

	var b interface{}
	bi := b.(int) // ERROR "type assertion must be checked"
	fmt.Println(bi)
}
