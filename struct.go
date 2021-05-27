package main

import (
	"fmt"
)

func main() {

	//
	// Struct
	//

	// 1
	var a struct {
		Name string
	} = struct{ Name string }{
		Name: "reza",
	}
	fmt.Println(a.Name)

	// 2
	c := struct {
		Name string
	}{
		"reza",
	}
	fmt.Println(c.Name)

	// 3
	var d struct {
		Name string
	}
	d.Name = "sina"
	fmt.Println(d.Name)

	// 4
	type aStruct struct {
		Name string
	}
	var f struct {
		aStruct
		Family string
	}
	f.Name = "salim"
	f.Family = "moazen"
	fmt.Println(f.Name)

	// 5
	var user struct {
		username string
		token    struct {
			value    string
			expireAt string
		}
	}
	user.username = "ali"
	user.token.value = "JWT sdfhighf"
	user.token.expireAt = "2019-10-10T00:00:00"
	fmt.Println(user.token.value)

	//
	// Dynamic Schema
	// https://medium.com/random-go-tips/dynamic-json-schemas-part-1-8f7d103ace71
	//
	var h map[string]interface{}
	h["name"] = "ali"
	fmt.Sprintf("%T", h)
	fmt.Println(h["name"])

	//
	// struct + array
	//
	var e = []struct {
		Name string
	}{
		{"Gholam"},
		{"Reza"},
	}
	fmt.Println(e[0].Name)

	//
	// struct in function arguments
	//
	testFn := func(class struct{ Name string }) {
		fmt.Println(class.Name)
	}
	g := struct{ Name string }{"salam"}
	testFn(g)

}
