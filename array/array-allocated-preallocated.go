package main

import "fmt"

func main() {

	var preAllocated, dynamic []int

	preAllocated = make([]int, 0, 5)
	dynamic = make([]int, 0)

	for i := 0; i < 5; i++ {
		preAllocated = append(preAllocated, i)
		dynamic = append(dynamic, i)
	}

	fmt.Println("Pre-allocated Slice:", preAllocated)
	fmt.Println("Dynamically Allocated Slice:", dynamic)
}
