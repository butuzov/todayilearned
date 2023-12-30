package main

import (
	"fmt"
	"log"
	"time"
)

func scheduler(term <-chan struct{}, waiter chan<- struct{}, workers chan struct{}, manager chan int) {
	for {
		select {
		case <-term:
			log.Println("Terminating manager execution")
			waiter <- struct{}{}
			return
		case i := <-manager:
			fmt.Println("Workers working b4 one start", len(workers))
			workers <- struct{}{}
			go func(n int) {

				fmt.Println("Workers working after job#", n, " done", len(workers))
				<-workers
				fmt.Println("Worker released")
			}(i)

		default:
			time.Sleep(25 * time.Microsecond)
		}
	}
}

func main() {
	var (
		limit   = 5
		stoper  = make(chan struct{})
		waiter  = make(chan struct{})
		worker  = make(chan struct{}, 3)
		manager = make(chan int)
	)

	go scheduler(stoper, waiter, worker, manager)

	for i := 0; i < limit; i++ {
		manager <- i
	}

	<-waiter
}
