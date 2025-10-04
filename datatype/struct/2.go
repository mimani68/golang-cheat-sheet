package main

import "fmt"

type Armor struct {
	Rate int
}

type Shield struct{}

// Post oop in golang
type Human struct {
	Armor
	Shield
	Name string
}

func (h *Human) SayHello() {
	fmt.Println("Hello")
	fmt.Println(h.Armor.Rate)
}

func main() {
	a := Human{}
	a.SayHello()
}
