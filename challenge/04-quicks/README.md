Write golang program which Create a recursive function (say quicksort()) to implement the quicksort. Partition the range to be sorted (initially the range is from 0 to N-1) and return the correct position of the pivot (say pi). Select the rightmost value of the range to be the pivot. Iterate from the left and compare the element with the pivot and perform the partition as shown above. Return the correct position of the pivot. Recursively call the quicksort for the left and the right part of the pi.


package main

import "fmt"

func Quicksort(arr []int, lo, hi int) {
	if lo < hi {
		p := partition(arr, lo, hi)
		Quicksort(arr, lo, p-1)
		Quicksort(arr, p+1, hi)
	}
}

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	i := lo - 1

	for j := lo; j <= hi-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[hi] = arr[hi], arr[i+1]
	return i + 1
}

func main() {
	arr := []int{4, 2, 6, 1, 9, 5}
	Quicksort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
