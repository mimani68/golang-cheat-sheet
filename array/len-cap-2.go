package main

import "fmt"

func main() {
	// Create an empty slice with a capacity of 5
	s := make([]int, 0, 5)
	fmt.Printf("Length: %d, Capacity: %d, Data: %v\n", len(s), cap(s), s)

	// OUTPUT
	// Length: 0, Capacity: 5, Data: []

	// Append elements until the capacity is reached
	for i := 0; i < 5; i++ {
		s = append(s, i)
		fmt.Printf("Length: %d, Capacity: %d, Data: %v\n", len(s), cap(s), s)

		// OUTPUT
		// Length: 1, Capacity: 5, Data: [0]
		// Length: 2, Capacity: 5, Data: [0 1]
		// Length: 3, Capacity: 5, Data: [0 1 2]
		// Length: 4, Capacity: 5, Data: [0 1 2 3]
		// Length: 5, Capacity: 5, Data: [0 1 2 3 4]
	}

	// Append one more element, triggering reallocation
	s = append(s, 10)
	fmt.Printf("Length: %d, Capacity: %d, Data: %v\n", len(s), cap(s), s)

	// OUTPUT
	// Length: 6, Capacity: 10, Data: [0 1 2 3 4 10]
}
