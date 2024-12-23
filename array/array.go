package main

import (
	"fmt"
	"math/rand"
)

//	Slices, on the other hand, are much more flexible, powerful, and
//
// convenient than arrays. Unlike arrays, slices can be resized using
// the built-in append function. Further, slices are reference types,
// meaning that they are cheap to assign and can be passed to other
// functions without having to create a new copy of its underlying array.
// Lastly, the functions in Go’s standard library all use slices rather than
// arrays in their public APIs.
//
// make([]Type, length, capacity)
// make([]Type, length)
// []Type{}
// []Type{value1, value2, ..., valueN}
func main() {
	var a = []int{1, 5, 6, 0, 4}
	fmt.Printf("a) Array: %v, Length: %d, Capacity: %d\n", a, len(a), cap(a))

	var b []int
	fmt.Printf("b) Array: %v, Length: %d, Capacity: %d\n", b, len(b), cap(b))

	c := []int{}
	fmt.Println(c)

	// Using make()
	d := make([]int, 0)
	e := append(d, 2, 4)
	fmt.Printf("d) Array: %v, Length: %d, Capacity: %d\n", d, len(d), cap(d))
	fmt.Printf("e) Array: %v, Length: %d, Capacity: %d\n", e, len(e), cap(e))

	// Using new
	f := new([]int)
	*f = append(*f, 1, 4, 4, 3)
	fmt.Printf("f) Array: %v, Length: %d, Capacity: %d\n", f, len(*f), cap(*f))

	// Filling array with random values
	g := make([]string, 2)
	for i := range g {
		g[i] = fmt.Sprintf("%d", rand.Intn(100))
	}

	// Append
	h := []string{}
	h = append(h, "one")
	h = append(h, "two")

	// Prepend
	i := append([]string{"zero"}, h...)
	fmt.Println(i)

	// Replace
	j := []int{3}
	j[0] = 1
	fmt.Println(j)

	// Access to item in list of strings
	k := []string{"ali", "mahdi"}
	fmt.Println(k[0])

	// Exclude one item in middle of array
	// Remove an item from list
	var brands []string
	brands = append(brands, "BMW", "HYNDAY", "KOWASAKI", "JET", "BOING")
	newBrands := append(brands[:2], brands[3:]...)
	fmt.Println(newBrands)

	// List of Type/User
	type User struct {
		Name string
	}
	var user = make([]User, 1) // instantiate and create a User object
	user[0].Name = "Hey"
	fmt.Printf("%T\n", user)
	fmt.Println(user[0])

}
