package main

import "fmt"

// Managing Slice Capacity to Avoid Memory Leaks

func main() {
	// Create a large slice
	largeSlice := make([]byte, 1<<20) // 1MB slice
	smallSlice := largeSlice[:10]     // 10-byte sub-slice

	// The large slice is still in memory because of the reference from smallSlice
	fmt.Printf("Length of smallSlice: %d, Capacity: %d\n", len(smallSlice), cap(smallSlice))

	// Copy only the necessary data to a new slice to release the large array
	newSmallSlice := make([]byte, 10)
	newSmallSlice = largeSlice[:10]
	fmt.Printf("Length of newSmallSlice: %d, Capacity: %d\n", len(newSmallSlice), cap(newSmallSlice))

	newSmallSlice = make([]byte, 10)
	copy(newSmallSlice, largeSlice[:10])
	fmt.Printf("Using copy) Length of newSmallSlice: %d, Capacity: %d\n", len(newSmallSlice), cap(newSmallSlice))

	largeSlice = nil // Remove the reference to the large slice

}
