package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func changeName(p *Person) {
	p.firstName = "Sina"
}

func main() {
	person := Person{
		firstName: "Ali",
		lastName:  "Kazemi",
	}

	fmt.Printf("%T\n", person)

	changeName(&person)

	fmt.Println(person)
}
