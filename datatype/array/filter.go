package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	// Create initial list
	list := []float32{}
	list = append(list, 1.3, 2.4, 3.0)

	// Example 1: Filter values greater than 2.0
	greaterThanTwo := slices.Collect(filter(list, func(v float32) bool {
		return v > 2.0
	}))
	fmt.Println("Values > 2.0:", greaterThanTwo)

	// Example 2: Filter values less than or equal to 2.4
	lessThanOrEqual := slices.Collect(filter(list, func(v float32) bool {
		return v <= 2.4
	}))
	fmt.Println("Values <= 2.4:", lessThanOrEqual)

	// Example 3: Filter exact value 3.0
	exactThree := slices.Collect(filter(list, func(v float32) bool {
		return v == 3.0
	}))
	fmt.Println("Values == 3.0:", exactThree)
}

// filter creates a sequence iterator from a slice with a predicate function
func filter[T any](s []T, predicate func(T) bool) iter.Seq[T] {
	return func(yield func(T) bool) {
		for _, v := range s {
			if predicate(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}
