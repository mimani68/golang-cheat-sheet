package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
}

func main() {
	person := Person{
		firstName: "Joe",
		lastName:  "Bloggs",
	}

	printIfPerson(person)
}

func printIfPerson(object interface{}) {
	person, ok := object.(Person)

	if ok {
		fmt.Printf("Hello %s!\n", person.lastName)
	}
}
