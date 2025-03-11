package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go initDb(&wg)
	go initServer(&wg)
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

// If wg.Add(3)
// Output will be
//
// CacheDb has started
// Server has started
// DB has started

// If wg.Add(2)
// Output will be
//
// CacheDb has started
// Server has started
