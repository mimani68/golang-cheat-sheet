package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	for i := 0; i < 3; i++ {
		go func() {
			fmt.Println(time.Now().Local().UTC().Format(time.RFC3339Nano))
		}()
	}
	time.Sleep(time.Second * 2)
}
