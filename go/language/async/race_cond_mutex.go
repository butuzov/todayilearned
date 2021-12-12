package main

import (
	"fmt"
	"sync"
)

func main() {

	var (
		counters = map[int]int{}
		mutex    = &sync.Mutex{}
		limit    = 5
		waiter   = &sync.WaitGroup{}
	)

	waiter.Add(limit)
	for i := 0; i < limit; i++ {
		go func(th int) {
			defer waiter.Done()
			for j := 0; j < limit; j++ {
				mutex.Lock()
				counters[th*10+j]++
				mutex.Unlock()
			}
		}(i)
	}
	waiter.Wait()
	fmt.Println("counters result", counters)

}
