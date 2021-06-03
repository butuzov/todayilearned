package main

import (
	"fmt"
	"sync"
)

// This example a kinda stupid, but lets assume i am lazy and there suppose to be
// heavy allocation of the instance "type" object, so we can get a ready one from
// the pool instead of creating new onces instead onces killed by gc.

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}
	myPool.Get()
	instance := myPool.Get()
	myPool.Put(instance)
	myPool.Get()

}
