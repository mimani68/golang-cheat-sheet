package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

func main() {

	messages := make(chan string)

	go func() {
		messages <- "ping"
		time.Sleep(time.Second)
		fmt.Println("Finish groutine")
	}()

	msg := <-messages
	fmt.Println(msg)

	//
	// Using channel for synchronization
	//
	done := make(chan bool, 1)
	go worker(done)

	<-done

}
