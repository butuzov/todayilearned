package main

import (
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	s, b := "string", []byte("string")

	spew.Dump(*(*[]byte)(unsafe.Pointer(&s)))
	spew.Dump(*(*string)(unsafe.Pointer(&b)))
}
