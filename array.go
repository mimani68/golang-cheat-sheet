package main

import "fmt"

func ArrayFn() {
	// 01 Define variable
	var a = []int{1, 5, 6, 0, 4}
	var b []int
	fmt.Println(a)
	fmt.Println(b)

	// 02 Dynamic array sized
	c := []int{}
	fmt.Println(c)

	// 03 Use make()
	d := make([]int, 0)
	c = append(d, 2, 4)
	d[0] = 1 // replace 1 -> 2
	fmt.Println(d)

	// 04 Append
	e := []string{}
	e = append(e, "one")
	e = append(e, "two")

	// 05 Preappend
	e = append([]string{"zero"}, e...)
	fmt.Println(e)

	// 05
	f := [1]string{}
	f[0] = "mahdi"
	fmt.Println(f[0])

	g := []string{"ali", "mahdi"}
	fmt.Println(g[0])

	// 06 slice
	h := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(h[1:4])
}
