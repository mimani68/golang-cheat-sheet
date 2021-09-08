package main

import (
	"fmt"
)

func StructFunction() {

	// 1
	var a struct {
		Name string
	} = struct{ Name string }{
		Name: "reza",
	}
	fmt.Println(a.Name)

	// 2
	b := struct {
		Name string
	}{
		"reza",
	}
	fmt.Println(b.Name)

	// 3
	var c struct {
		Name string
	}
	c.Name = "sina"
	fmt.Println(c.Name)

	// 4
	type aStruct struct {
		Name string
	}
	var d struct {
		aStruct
		Family string
	}
	d.Name = "salim"
	d.Family = "moazen"
	fmt.Println(d.Name)

	// 5
	var e struct {
		username string
		token    struct {
			value    string
			expireAt string
		}
	}
	e.username = "ali"
	e.token.value = "JWT sdfhighf"
	e.token.expireAt = "2019-10-10T00:00:00"
	fmt.Println(e.token.value)

	f := struct {
		Name    string
		Account struct {
			Username string
			Password string
		}
	}{
		Name: "ali",
		Account: struct {
			Username string
			Password string
		}{
			Username: "mimani",
			Password: "123",
		},
	}
	fmt.Println(f)

	//
	// inline struct handeling
	//
	fmt.Println(struct {
		Name string
	}{
		Name: "mahdi",
	})

	//
	// struct + array
	//
	var g = []struct {
		Name string
	}{
		{"Gholam"},
		{"Reza"},
	}
	fmt.Println(g[0].Name)

	//
	// struct in function arguments
	//
	testFn := func(class struct{ Name string }) {
		fmt.Println(class.Name)
	}
	h := struct{ Name string }{"salam"}
	testFn(h)

}
