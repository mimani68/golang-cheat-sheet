package main

import (
	"fmt"
)

func main() {

	//
	// ⛔ This code won't work
	//
	// var h map[string]interface{}
	// h["name"] = "ali"
	// fmt.Sprintf("%T", h)
	// fmt.Println(h["name"])

	//
	// map[string]interface{}
	//
	a := map[string]interface{}{
		"bacon": "delicious",
		"eggs": struct {
			source string
			price  float64
		}{"chicken", 1.75},
		"steak": true,
	}
	a["bacon"] = "bad"
	fmt.Println(a["bacon"])

	//
	// Generate a map[string]interface{}
	//
	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = 23
	myData["attribute"] = map[string]int{
		"strength":     100,
		"agility":      100,
		"intelligence": 100,
	}
	print(myData)

	//
	// Direct definition
	//
	fmt.Println(map[string]interface{}{
		"title": "JWT 接口测试",
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiaWQiOiIxIiwibmFtZSI6IkhvYm8ifQ.YUzBykoELyKoQWaugkVNf3d09HBhICBJoOcWQKnveRQ",
	})

}
