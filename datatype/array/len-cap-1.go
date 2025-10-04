package main

import "fmt"

func main() {
	// Create a slice with a specified length and capacity
	s := make([]int, 3, 6)
	fmt.Printf("Length: %d, Capacity: %d, Data: %v\n", len(s), cap(s), s)

	// Output:
	// Length: 3, Capacity: 6, Data: [0 0 0]

	// Append elements to the slice
	s = append(s, 1, 2, 3)
	fmt.Printf("Length: %d, Capacity: %d, Data: %v\n", len(s), cap(s), s)

	// Output:
	// Length: 6, Capacity: 6, Data: [0 0 0 1 2 3]

	// If we append more elements than the capacity, Go will allocate a new underlying array
	s = append(s, 4)
	fmt.Printf("Length: %d, Capacity: %d, Data: %v\n", len(s), cap(s), s)

	// Output:
	// Length: 7, Capacity: 12, Data: [0 0 0 1 2 3 4]
}
