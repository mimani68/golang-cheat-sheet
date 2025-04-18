package main

import "fmt"

var key = "name"

func main() {
	list := []interface{}{"first"}
	list = append(list, string("second"))

	Application(list)
}

func Application(list []interface{}) {
	user := map[string]interface{}{}
	user[key] = "ali"
	list = append(list, user)

	userTwo := map[string]interface{}{}
	userTwo[key] = "sajad"
	list = append(list, userTwo)

	userThree := struct {
		Name   string
		Family string
	}{
		Name:   "Mahdi",
		Family: "Imani",
	}
	list = append(list, userThree)

	fmt.Println(list)
	fmt.Println(list...)

	printInterfaceName(list)
}

func printInterfaceName(list []interface{}) {
	for index, _ := range list {
		fmt.Println(list[index])
	}
}
