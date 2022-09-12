package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	c := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for {
		select {
		case <-c:
			fmt.Println("Break the loop")
			time.Sleep(time.Second * 2)
			fmt.Println("BYE")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Hello in a loop")
		}
	}
}
