package main

import (
	"unsafe"
	"fmt"
)

// example from the Aliaksand Valialkin presentation about unsafe.
// pointless, but ok. not really bce, but it is in some order.

func example(){
	a := []int{1,2,3}
	x := unsafeGet(a, 5)

	fmt.Printf("x:%+v\n", x)
}

func unsafeGet(a []int, idx int) int {
	// calculating offset from zero element for index idx with size of zero element.
	return *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&a[0])) + uintptr(idx)*unsafe.Sizeof(a[0])))
}
