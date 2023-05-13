// https://www.interviewbit.com/problems/maximum-ones-after-modification/
// Write golang program which Given a binary array  A := []int{1, 0, 0, 1, 1, 0, 1} and a number  B := 1

package main

import "fmt"

func longestSubSegment(A []int, B int) int {
	left, right := 0, 0
	length := len(A)
	maxLen, zeroCount := 0, 0

	for right < length {
		if A[right] == 0 {
			zeroCount++
		}
		for zeroCount > B {
			if A[left] == 0 {
				zeroCount--
			}
			left++
		}
		maxLen = max(maxLen, right-left+1)
		right++
	}

	return maxLen
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	A := []int{1, 0, 0, 1, 1, 0, 1}
	B := 2
	result := longestSubSegment(A, B)
	fmt.Println("Longest subsegment of 1s:", result)
}
