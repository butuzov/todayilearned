package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
	Difference on blocking and non blocking counters.
	Atomic operations.
*/

func main() {

	var (
		simpleCounter     int32 = 0
		atomicCounter     int32 = 0
		atomicCounterFunc       = func() { atomic.AddInt32(&atomicCounter, 1) }
		simpleCounterFunc       = func() { simpleCounter++ }
	)

	for i := 0; i < 100000; i++ {
		go atomicCounterFunc()
		go simpleCounterFunc()
	}

	time.Sleep(2 * time.Millisecond)
	fmt.Println("Results: Atomic", simpleCounter)
	fmt.Println("Results: Simple", atomicCounter)
}
