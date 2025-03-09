package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	// wg.Done()
}

func main() {
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		go increment(&m)
	}
	time.Sleep(time.Second * 3)
	fmt.Println("final value of x", x)

	// var w sync.WaitGroup
	// var m sync.Mutex
	// for i := 0; i < 1000; i++ {
	// 	w.Add(1)
	// 	go increment(&w, &m)
	// }
	// w.Wait()
	// fmt.Println("final value of x", x)
}
