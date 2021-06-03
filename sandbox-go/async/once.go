package main

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func main() {

	runtime.GOMAXPROCS(0)

	var limit = 5
	var t = struct {
		sync.Once
		sync.WaitGroup
		Winner int
	}{}

	t.Add(limit)
	for i := 1; i <= limit; i++ {
		go func(n int) {

			defer t.Done()
			time.Sleep(35 * time.Microsecond)

			t.Do(func() {
				t.Winner = n
				fmt.Println("Only", strconv.Itoa(t.Winner)+"th", "came to the finish")
			})
		}(i)
	}

	t.Wait()
}
