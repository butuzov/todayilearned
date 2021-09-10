package testdata

import (
	"fmt"
	"runtime"
)

func GC() {
}

func foogc() {
	fmt.Println("just testing")
	GC()
	runtime.Goexit()
	runtime.GC() // MATCH /explicit call to the garbage collector/
}
