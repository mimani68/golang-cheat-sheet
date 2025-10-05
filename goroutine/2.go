package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func(ch chan string) {
		time.Sleep(time.Second * time.Duration(1))
		ch <- "ping"
	}(channel)

	fmt.Println(<-channel)

	// Output
	// ping

}
