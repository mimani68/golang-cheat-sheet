package main

import "fmt"

func main() {
	/* a slice of unspecified size */
	var numbers []int
	numbers = []int{0, 0, 0, 0, 0}

	/* a slice of length 5 and capacity 5*/
	numbers = make([]int, 5, 5)

	/* missing lower bound implies 0*/
	fmt.Println(numbers)
	fmt.Println(numbers[:3])

	/* add one element to slice*/
	numbers = append(numbers, 1)

	/* copy content of numbers to numbers1 */
	numbers1 := []int{1, 2}
	copy(numbers, numbers1)
	fmt.Println(numbers)
}
