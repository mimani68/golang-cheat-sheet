package main

import (
	"fmt"
	"sync"
	"time"
)

// A WaitGroup waits for a collection of goroutines to finish executing
var wg sync.WaitGroup

func main() {
	for i := 0; i < 3; i++ {
		// This counter represents the number of goroutines that the WaitGroup is waiting for.
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(time.Now().Local().UTC().Format(time.RFC3339Nano))
		}()
	}
	// `wg.Wait()` or `time.Sleep(time.Second * 2)` are acting interchangeably.
	time.Sleep(time.Second * 2)
}

// Outout
//
// 2025-03-11T16:02:37.079893472Z
// 2025-03-11T16:02:37.079941921Z
// 2025-03-11T16:02:37.07992148Z
