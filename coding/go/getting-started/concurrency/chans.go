package main

import "fmt"

func main() {
	// buffered
	cb := make(chan int, 1)
	go func() { fmt.Println(<-cb) }()
	cb <- 1

	// unbuffered
	cu := make(chan int)
	go func() { fmt.Println(<-cu) }()
	cu <- 2
}
