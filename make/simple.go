package main

import "fmt"

func main() {
	// The "Zero Length, High Capacity" Pattern
	s1 := make([]int, 0, 5)
	fmt.Println(s1)
	// []

	s2 := make([]int, 2, 5)
	fmt.Println(s2)
	// [0 0]

	// The "Pre-filled" Pattern
	// We do NOT need to use append.
	s3 := make([]int, 5)
	fmt.Println(s3)
	// [0 0 0 0 0 0]

	s4 := make([]int, 5)
	s4 = append(s4, 0)
	fmt.Println(s4)
	// [0 0 0 0 0 0 0]

	s5 := make([]int, 10)
	// Create a 'child' slice from index 3 to 6 (exclusive)
	s5_child := s5[3:6]
	fmt.Println(s5_child)
	fmt.Printf("Capacity=%d Length=%d\n", cap(s5_child), len(s5_child))
	// Length = 3 (elements 4, 5, 6)
	// Capacity = 7 (from index 3 to the end of the parent array)
}
