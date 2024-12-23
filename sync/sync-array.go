package main

import (
	"fmt"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return make([]int, 0, 100)
	},
}

func main() {
	// Get a slice from the pool
	s := pool.Get().([]int)
	defer pool.Put(s[:0]) // Reset the slice and return it to the pool

	// Use the slice
	for i := 0; i < 100; i++ {
		s = append(s, i)
	}
	fmt.Printf("Length: %d, Capacity: %d\n", len(s), cap(s))
}
