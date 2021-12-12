package main

import (
	"fmt"
)

func main() {
	in := make(chan int, 0)
	go func(out chan<- int) {
		for i := 0; i <= 10; i++ {
			out <- i
		}
		close(out)
	}(in)

	for i := range in {
		fmt.Println(i)
	}
}
