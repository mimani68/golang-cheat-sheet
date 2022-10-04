package main

import "fmt"

type Armor struct{}

type Shield struct{}

type Human struct {
	Armor
	Shield
	Name string
}

func (h *Human) SayHello() {
	fmt.Println("Hello")
	fmt.Println(h.Armor.xxxx)
}

func main() {
	a := Human{}
	a.SayHello()
}
