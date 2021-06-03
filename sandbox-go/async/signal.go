package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var (
		waiter = make(chan struct{})
		notify = make(chan os.Signal)
	)
	signal.Notify(notify, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		fmt.Println("To win the game you must kill me, John Romero")
		fmt.Println(<-notify)
		fmt.Println("Finish him!")
		fmt.Println(<-notify)
		waiter <- struct{}{}
	}()

	<-waiter
}
