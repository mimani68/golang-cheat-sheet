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

	channel := make(chan string)

	go func() {
		time.Sleep(time.Second * time.Duration(3))
		channel <- "ping"
	}()

	msg := <-channel
	fmt.Println(msg)

	//
	// Using channel for synchronization
	//
	// done := make(chan bool, 1)
	// go worker(done)

	// <-done

}
