package main

import (
	"fmt"
)

func main() {
	chTerm := make(chan struct{})
	chData := make(chan int)

	// Generator
	go func(chTerm <-chan struct{}, dataCh chan<- int) {
		var val = 0
		for {
			select {
			case <-chTerm:
				return
			case dataCh <- val:
				val++
			}
		}
	}(chTerm, chData)

	// Generator Consumer
	for curVal := range chData {
		fmt.Println("got", curVal)
		if curVal >= 3 {
			fmt.Println("stop it!")
			chTerm <- struct{}{}
			break
		}
	}

}
