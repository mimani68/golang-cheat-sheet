package main

import "fmt"

func main() {

	// Why pre‑allocate?
	// If you know roughly how many elements you’ll add, pre‑allocating capacity reduces the number of reallocations and copies, which can improve performance in tight loops or large data‑processing tasks.
	var preAllocated, dynamic_a, dynamic_b []int

	// 1️⃣  Create a slice with an "initial length" of 3 and capacity of 5.
	// make([]T, length, capacity)
	//   Length: number of elements the slice currently holds.
	//   Capacity: size of the underlying array; the maximum number of elements it can hold before a new allocation is needed.
	preAllocated = make([]int, 3, 5)

	// 2️⃣  Create an empty slice (length 0, capacity 0) that will grow as needed.
	// dynamic = make([]int) // invalid operation: make([]int) expects 2 or 3 arguments; found 1
	dynamic_a = make([]int, 0)
	dynamic_b = make([]int, 5)

	// 3️⃣  Append five values (0‑4) to each slice.
	for i := 0; i < 5; i++ {
		preAllocated = append(preAllocated, i)
		dynamic_a = append(dynamic_a, i)
		dynamic_b = append(dynamic_b, i)
	}

	// 4️⃣  Print the final contents of both slices.
	fmt.Println("Pre-allocated Slice with initial values:", preAllocated)
	// Pre-allocated Slice: [0 0 0 0 1 2 3 4]

	fmt.Println("Dynamically Allocated Slice:", dynamic_a)
	// Dynamically Allocated Slice: [0 1 2 3 4]

	fmt.Println("Dynamically Allocated Slice with initial values:", dynamic_b)
	// Dynamically Allocated Slice with initial values: [0 0 0 0 0 0 1 2 3 4]
}
