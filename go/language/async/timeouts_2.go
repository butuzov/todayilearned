package main

import (
	"fmt"
	"math/rand"
	"time"
)

func times() []time.Duration {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	var to = []time.Duration{time.Second, 2 * time.Second, 3 * time.Second}

	r.Shuffle(len(to), func(i, j int) { to[i], to[j] = to[j], to[i] })
	return to
}

func slowpoke(t time.Duration) chan struct{} {
	ch := make(chan struct{})
	go func() {
		time.AfterFunc(t, func() { ch <- struct{}{} })
	}()
	return ch
}

func main() {
	var timeouts = times()
	timer := time.NewTimer(timeouts[0])
	select {
	case <-timer.C:
		fmt.Println("Timeout happend")
	case <-time.After(timeouts[1]):
		fmt.Println("time.After timeout happend")
	case <-slowpoke(timeouts[2]):
		if !timer.Stop() {
			<-timer.C
		}
		fmt.Println("Slow func is done")
	}
}
