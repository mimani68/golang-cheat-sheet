package main

import "fmt"

func main() {
	a := map[string]string{}
	a["id"] = "id094u598"
	a["username"] = "mimani"
	privateFunc(map[string]string{
		"id": "id2384569",
	})
}

func privateFunc(a ...interface{}) {
	fmt.Println(a)
	fmt.Println(a...)
	for _, param := range a {
		fmt.Println(param)
	}
}
