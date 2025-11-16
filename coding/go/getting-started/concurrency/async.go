package main

import (
	"fmt"
	"sync"
)

// ----------------
// This is simple "inheritance" via embeding, but beware of such implementations.
// It exposes public methods of embeded structs, as sideeffect.
// ----------------

type AsyncStruct struct {
	sync.Once
	sync.Mutex
	sync.WaitGroup

	value int
}

func (a *AsyncStruct) Increase() {
	a.Add(1)

	go func() {
		a.Lock()
		defer a.Unlock()
		defer a.Done()

		a.value++
	}()
}

func main() {
	a := AsyncStruct{}
	a.Increase()
	a.Wait()

	fmt.Print(a.value)
}
