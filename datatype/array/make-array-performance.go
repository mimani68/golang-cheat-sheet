package main

import (
	"fmt"
	"time"
)

func main() {
	// Measure performance with pre-allocation
	start := time.Now()
	preAllocated := make([]int, 0, 10000)
	for i := 0; i < 10000; i++ {
		preAllocated = append(preAllocated, i)
	}
	fmt.Printf("Pre-allocated time: %v\n", time.Since(start))
	fmt.Printf("Pre-allocated Length: %d, Capacity: %d\n", len(preAllocated), cap(preAllocated))

	// OUTPUT
	// Pre-allocated time: 76.577µs
	// Pre-allocated Length: 10000, Capacity: 10000

	// Measure performance without pre-allocation
	start = time.Now()
	dynamic := make([]int, 0)
	for i := 0; i < 10000; i++ {
		dynamic = append(dynamic, i)
	}
	fmt.Printf("Dynamic allocation time: %v\n", time.Since(start))
	fmt.Printf("Dynamic Length: %d, Capacity: %d\n", len(dynamic), cap(dynamic))

	// OUTPUT
	// Dynamic allocation time: 345.614µs
	// Dynamic Length: 10000, Capacity: 12288
}
