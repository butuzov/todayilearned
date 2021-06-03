package main

import (
	"context"
	"log"
	"math/rand"
	"time"
)

// There are some chances that function will succeed and some changes that not.

func main() {

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(rand.Int63n(7))*time.Second)
	defer func() {
		log.Printf("ctx cancel:  %-10s", time.Since(start))
		cancel()
	}()

	Runner(ctx, start)
}

func Runner(ctx context.Context, start time.Time) {
	var res = make(chan struct{})
	go func() {
		LongRunningFunction(ctx, start)
		res <- struct{}{}
	}()

	select {
	case <-res:
		log.Printf("func done:   %-10s", time.Since(start))
	case <-ctx.Done():
		log.Printf("ctx timeout: %s", ctx.Err())
		log.Printf("ctx timeout: %-10s", time.Since(start))
	}

}

func LongRunningFunction(ctx context.Context, start time.Time) {
	time.Sleep(3 * time.Second)
}
