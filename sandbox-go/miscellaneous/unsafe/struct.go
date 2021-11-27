package main

import (
	"unsafe"

	"github.com/davecgh/go-spew/spew"
)

// packing struct to butes

func main() {
	// structs...
	type Struct struct {
		Src     int32
		Dst     int32
		SrcPort uint16
		DstPort uint16
	}

	var struct_value Struct = Struct{1, 2, 3, 4}
	const size = int(unsafe.Sizeof(Struct{}))
	var asByteSlice []byte = (*(*[size]byte)(unsafe.Pointer(&struct_value)))[:]

	spew.Dump(asByteSlice)
}
