package main

import (
	"fmt"
	"math/rand"
	"time"
)

func times() []time.Duration {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var to = []time.Duration{time.Second, 2 * time.Second}

	r.Shuffle(len(to), func(i, j int) { to[i], to[j] = to[j], to[i] })
	return to
}

func main() {

	var timeouts = times()

	var (
		start   = time.Now()
		ender   = make(chan struct{})
		timeout = time.After(timeouts[0])
		timer   = time.AfterFunc(timeouts[1], func() {
			fmt.Printf("Hello World took %v to execute\n", time.Now().Sub(start))
			ender <- struct{}{}
		})
	)

	for {
		select {
		case <-ender:
			fmt.Println("Execution terminated by signal routine")
			return
		case <-timeout:
			fmt.Printf("Timeout trigerred? (%t)\n", timer.Stop())
			return
		default:
			time.Sleep(25 * time.Microsecond)
		}
	}

	time.Sleep(time.Second)
}
