package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	// ------------------------------------------------------------------------
	// this is why go need generics

	nums := func(ctx context.Context) <-chan int {
		var (
			dst = make(chan int)
			num = 1
		)

		go func() {
			for {
				select {
				case <-ctx.Done():
					log.Println("exit generator")
					return
				case dst <- num:
					num++
				}
			}
		}()

		return dst
	}

	tick := func(ctx context.Context) <-chan time.Time {
		var dst = make(chan time.Time)
		go func() {
			for {
				select {
				case dst <- time.Now():
				case <-ctx.Done():
					log.Println("exit timer")
					return
				}
			}
		}()
		return dst
	}

	// ------------------------------------------------------------------------
	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	var (
		intsChan = nums(ctx)
		timeChan = tick(ctx)
		start    = time.Now()
	)

	// Onces cancel called for context with Cancel
	// its proragated to ctx.Done for every context receiver

	for {
		select {
		case n := <-intsChan:
			fmt.Println(n)
			if n >= 10 {
				cancel()
				time.Sleep(time.Second)
				return
			}
		case t := <-timeChan:
			fmt.Println(t.Sub(start))
		}
	}

}
