package main

import "fmt"

func main() {

	a := "Hey"
	fmt.Println("show value: ", a)                                 // "Hey"
	fmt.Println("show memory addess of varable: ", &a)             // 0xc000010240
	fmt.Println("find memory address, then return value: ", *(&a)) // "Hey"

	b := &struct {
		Name string
	}{
		Name: "Ali",
	}
	fmt.Println("show value: ", b)                                 // {Ali}
	fmt.Println("show memory addess of varable: ", &b)             // 0xc0000b6020
	fmt.Println("find memory address, then return value: ", *b)    // {Ali}
	fmt.Println("find memory address, then return value: ", *(&b)) // &{Ali}

}
