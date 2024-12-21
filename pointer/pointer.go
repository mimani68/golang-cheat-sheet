package main

import "fmt"

func main() {

	a := "Hey"
	fmt.Println("a")
	fmt.Println("show value: ", a)                                 // "Hey"
	fmt.Println("show memory address of variable: ", &a)           // 0xc000010240
	fmt.Println("find memory address, then return value: ", *(&a)) // "Hey"

	b := new(string)
	*b = "hey"
	fmt.Println("b")
	fmt.Println("show value: ", *b)                     // "hey"
	fmt.Println("show memory address of variable: ", b) // 0xc0000140b0
	fmt.Println("find memory address: ", &b)            // 0xc00005c048

	c := &struct {
		Name string
	}{
		Name: "Ali",
	}
	fmt.Println("c")
	fmt.Println("show value: ", c)                                 // {Ali}
	fmt.Println("show memory address of variable: ", &c)           // 0xc0000b6020
	fmt.Println("find memory address, then return value: ", *c)    // {Ali}
	fmt.Println("find memory address, then return value: ", *(&c)) // &{Ali}

}
