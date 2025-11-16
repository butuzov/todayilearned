package main

import (
	"context"
	"log"
	"math/rand"
	"sync/atomic"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {

	tot := int64(5)
	sem := semaphore.NewWeighted(tot) // Total Semaphores available.
	ctx := context.Background()

	var val int32

	for i := 0; i < 25; i++ {

		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(n int) {
			defer func() {
				sem.Release(1)
				atomic.AddInt32(&val, -1)
			}()

			time.Sleep(time.Duration(rand.Int63n(1)) * time.Millisecond)
			atomic.AddInt32(&val, 1)
			log.Printf("%d) Current Counter is %d\n", n, atomic.LoadInt32(&val))
		}(i)
	}
	// Check for released semaphores
	if !sem.TryAcquire(tot) {
		log.Printf("Still in progress")
	}

	// Check for released semaphores
	if err := sem.Acquire(ctx, tot); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	} else {
		log.Printf("We are done")
	}
}
