package main

import (
	"encoding/json"
	"fmt"
	"json"
)

type Person struct {
	firstName string
	lastName  string
}

func main() {

	//
	// convert struct to map[string]interface{}
	//
	b := struct {
		Id string
	}{
		Id: "w984y9n84y9r84",
	}
	fmt.Println(b.Id)
	var c map[string]interface{}
	inrec, _ := json.Marshal(b)
	json.Unmarshal(inrec, &c)

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
