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
	a1, b1 := Vertex{1, 2}, Vertex{5, 3}
	a1 = b1
	b1.X = 2
	fmt.Println(a1.X, a1.Y) // => 5 3
	fmt.Printf("%T\n", a1)  // => main.Vertex

	// Point to memory address
	a2, b2 := &Vertex{1, 2}, &Vertex{5, 3}
	a2 = b2
	b2.X = 15
	fmt.Println(a2.X, a2.Y) // => 15 3
	fmt.Printf("%T\n", a2)  // => *main.Vertex

	fmt.Println("--------")

	v3 := new(Vertex)
	v3.X = 2
	(*v3).Y = 4
	fmt.Printf("%T - %v\n", v3, v3)   // => *main.Vertex - &{2 4}
	fmt.Printf("%T - %v\n", &v3, &v3) // => **main.Vertex - 0xc00000e030
	fmt.Printf("%T\n", *v3)           // => main.Vertex
	fmt.Printf("%T\n", v3.Y)          // => int
	fmt.Printf("%T\n", v3.X)          // => int

}
