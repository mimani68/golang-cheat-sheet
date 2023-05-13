package main

import (
	"fmt"
	"time"
)

type CollectionsHolder[T any] struct {
	// type CollectionsHolder[T User | float32] struct {
	Value T
	Time  time.Time
}

type User struct {
	Name string
}

func main() {
	a := CollectionsHolder[string]{Value: "Hello, World!"}
	b := CollectionsHolder[int]{Value: 42}
	c := CollectionsHolder[float64]{Value: 3.14}
	d := CollectionsHolder[User]{Value: User{
		Name: "mahdi",
	}}

	fmt.Printf("a: %v\n", a.Value)
	fmt.Printf("b: %v\n", b.Value)
	fmt.Printf("c: %v\n", c.Value)
	fmt.Printf("c: %v\n", c.Value.Name)
	fmt.Printf("d: %v\n", d.Value)
	fmt.Printf("d: %v\n", d.Value.Name)
}
