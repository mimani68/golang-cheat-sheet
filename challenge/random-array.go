package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	nums := [10]int{}
	for i := range nums {
		nums[i] = rand.Intn(10) // generate random integer between 0 and 9
	}

	uniqueNums := make([]int, 0, len(nums)) // create an empty slice to hold the unique elements
	seen := make(map[int]bool)              // create a map to store the values we've seen

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			uniqueNums = append(uniqueNums, num)
		}
	}

	fmt.Println("Original array:", nums)
	fmt.Println("Unique values:", uniqueNums)
}
