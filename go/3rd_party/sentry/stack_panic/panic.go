package main

import "errors"

//go:noinline
func func1() {
	func2()
}

//go:noinline
func func2() {
	func3()
}

//go:noinline
func func3() {
	func4()
}

//go:noinline
func func4() {
	func5()
}

//go:noinline
func func5() {
	panic(errors.New("ooops"))
}
