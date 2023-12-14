package main

import "fmt"

func init() {
	fmt.Println("[BOOTSTRAP] Hello app")
}

type Human struct {
	Name string ``
}

func (h Human) SetName(name string) bool {
	h.Name = name
	return true
}

func main() {
	human := Human{}
	_ = human.SetName("ali")
	fmt.Printf("[VALUE] %v\n", human.Name)
}
