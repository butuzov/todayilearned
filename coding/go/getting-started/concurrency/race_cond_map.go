package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		mapa   = &sync.Map{}
		limit  = 5
		waiter = &sync.WaitGroup{}
	)

	waiter.Add(limit)
	for i := 0; i < limit; i++ {
		go func(th int) {
			defer waiter.Done()
			for j := 0; j < limit; j++ {
				mapa.LoadOrStore(th*10+j, 1)
			}
		}(i)
	}
	waiter.Wait()

	var counters = map[int]int{}
	mapa.Range(func(key, value interface{}) bool {
		counters[key.(int)] = value.(int)
		return true
	})
	fmt.Println("counters result", counters)

}
