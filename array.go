package main

import "fmt"

func ArrayFn() {
	// 01
	var a = []int{1, 5, 6, 0, 4}
	var b []int

	// 02 dynamic array sized
	c := []int{}
	fmt.Println(c)

	// 03
	c := make([]int, 0)
	c = append(c, 2, 4)
	c[0] = 1 // replace 1 -> 2
	fmt.Println(c)

	// 04
	e := [1]string{}
	e[0] = "mahdi"
	fmt.Println(e[0])

	f := []string{"ali", "mahdi"}
	fmt.Println(f[0])

	// 04 slice
	n := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(n[1:4])
}
