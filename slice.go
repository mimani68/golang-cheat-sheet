package main

import "fmt"

//
//  Slices, on the other hand, are much more flexible, powerful, and
// convenient than arrays. Unlike arrays, slices can be resized using
// the built-in append function. Further, slices are reference types,
// meaning that they are cheap to assign and can be passed to other
// functions without having to create a new copy of its underlying array.
// Lastly, the functions in Go’s standard library all use slices rather than
// arrays in their public APIs.
//
//
// make([]Type, length, capacity)
// make([]Type, length)
// []Type{}
// []Type{value1, value2, ..., valueN}
//
func main() {
	var a = []int{1, 5, 6, 0, 4}
	fmt.Println(a)

	var b []int
	fmt.Println(b)

	c := []int{}
	fmt.Println(c)

	// Using make()
	d := make([]int, 0)
	c = append(d, 2, 4)
	fmt.Println(d)

	// Append
	e := []string{}
	e = append(e, "one")
	e = append(e, "two")

	// Preappend
	e = append([]string{"zero"}, e...)
	fmt.Println(e)

	// Replace
	f := []int{3}
	f[0] = 1
	fmt.Println(f)

	// list of strings
	g := []string{"ali", "mahdi"}
	fmt.Println(g[0])

	// Parts of list
	h := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(h[1:4])
}