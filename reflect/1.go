package main

import (
	"crypto/rand"
	"fmt"
	"reflect"
)

type User struct{}

func main() {
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf("salam"))
	fmt.Println(reflect.TypeOf(rand.Int))
	fmt.Println(reflect.TypeOf(User{}))
}
