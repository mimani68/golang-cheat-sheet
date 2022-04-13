package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

func main() {
	// Point to value
	v := Vertex{1, 2}
	fmt.Println(v.X, v.Y)
	fmt.Printf("%T\n", v.X)

	// Point to memory address
	vv := new(Vertex)
	vv.X = 2
	(*vv).Y = 4
	fmt.Printf("%T\n", vv)   // => *main.Vertex
	fmt.Printf("%T\n", &vv)  // => **main.Vertex
	fmt.Printf("%T\n", *vv)  // => main.Vertex
	fmt.Printf("%T\n", vv.Y) // => int
}
