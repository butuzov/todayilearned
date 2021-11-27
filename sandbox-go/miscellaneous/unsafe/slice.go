package main

import (
	"reflect"
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	/// slices...
	data := []int{1, 2, 3, 4, 5}
	var out []byte
	in := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	outHeader := (*reflect.SliceHeader)(unsafe.Pointer(&out))
	outHeader.Data = in.Data
	outHeader.Len = in.Len * int(unsafe.Sizeof(&data[0]))
	outHeader.Cap = in.Cap * int(unsafe.Sizeof(&data[0]))

	// out is altered!

	spew.Dump(out)
}
