package main

import (
	"fmt"
	"sync"
)

// this is simple "inheritance" via embeding, but beware of such implementations.
// they are exposed publick methods of embeded structs, so not its a part of Your API.
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
	var a = AsyncStruct{}
	a.Increase()
	a.Wait()

	fmt.Print(a.value)
}
