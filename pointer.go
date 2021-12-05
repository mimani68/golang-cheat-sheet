package main

import "fmt"

func main() {

	a := "salam"
	fmt.Println("show value: ", a)                                 // "salam"
	fmt.Println("show memory addess of varable: ", &a)             // 0xc000010240
	fmt.Println("find memory address, then return value: ", *(&a)) // "salam"

}
