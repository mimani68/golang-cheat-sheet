package main

import "fmt"

func ArrayFn() {
	a := []int{} // dynamic array sized
	fmt.Println(">", a)

	var b []int

	c := make([]int, 0)
	c = append(b, 2, 4)
	c[0] = 1 // replace 1 -> 2
}
