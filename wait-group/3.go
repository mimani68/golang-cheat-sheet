package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go initDb(&wg)
	wg.Add(1)
	go initServer(&wg)
	wg.Add(1)
	go initCacheDb(&wg)
	wg.Wait()
}

func initDb(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("DB has started")
}

func initServer(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Server has started")
}

func initCacheDb(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("CacheDb has started")
}

//
// Output
//
// CacheDb has started
// Server has started
// DB has started
