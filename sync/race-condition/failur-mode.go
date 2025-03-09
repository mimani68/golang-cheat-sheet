package main

import (
	"fmt"
	"time"
)

var x = 0

func increment() {
	x = x + 1
}

func main() {
	for i := 0; i < 1000; i++ {
		go increment()
	}
	time.Sleep(time.Second * 3)
	fmt.Println("final value of x", x)
}
