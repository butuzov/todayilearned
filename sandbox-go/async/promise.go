package main

import (
	"fmt"
	"time"
)

// Promise like patterns, we returning a channel to get results of our goroutine.

func executor() chan struct{} {

	var result = make(chan struct{})

	go func(pipe chan<- struct{}) {
		time.Sleep(1 * time.Second)
		fmt.Println("Execution done")
		pipe <- struct{}{}
	}(result)

	return result
}

func main() {
	signal := executor()
	<-signal
}
