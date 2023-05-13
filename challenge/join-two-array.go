package main

import "fmt"

func mergeTwoLists(A []int, B []int) []int {
	i, j := 0, 0
	merged := []int{}

	for i < len(A) && j < len(B) {
		if A[i] < B[j] {
			merged = append(merged, A[i])
			i++
		} else {
			merged = append(merged, B[j])
			j++
		}
	}

	for i < len(A) {
		merged = append(merged, A[i])
		i++
	}

	for j < len(B) {
		merged = append(merged, B[j])
		j++
	}

	return merged
}

func main() {
	// create two sorted integer slices
	A := []int{1, 5, 8}
	B := []int{10, 23}

	// merge the slices and print the resulting slice
	merged := mergeTwoLists(A, B)
	fmt.Println(merged)
}
