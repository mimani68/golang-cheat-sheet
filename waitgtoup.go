package main

import (
	"fmt"
	"sync"
	"time"
)

func getBunchOfData(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Start fetch data")
	time.Sleep(time.Second * 2)
	fmt.Println("Data gathered from source")
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	defer wg.Wait()
	go getBunchOfData(wg)
	time.Sleep(time.Millisecond * 300)
	go getBunchOfData(wg)
}
