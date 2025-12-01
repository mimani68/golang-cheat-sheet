// Write golang program Given an array A := []int{-1, 2, 1, -4} of N integers, find three integers in A such that the sum is closest to a given number B := 1. Return the sum of those three integers.
// https://www.interviewbit.com/problems/3-sum/

package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2; i++ {
		left := i + 1
		right := n - 1
		for left < right {
			curSum := nums[i] + nums[left] + nums[right]
			if curSum > target {
				right--
			} else {
				left++
			}
			if math.Abs(float64(target-curSum)) < math.Abs(float64(target-closestSum)) {
				closestSum = curSum
			}
		}
	}
	return closestSum
}

func main() {
	// A := []int{-1, 2, 1, -4}
	// B := 1
	A := []int{1, 2, 3}
	B := 6
	result := threeSumClosest(A, B)
	fmt.Println(result)
}
