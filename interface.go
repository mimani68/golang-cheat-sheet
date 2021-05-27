package main

import (
	"encoding/json"
	"fmt"
)

func SampleInterface() {

	privateFunc(map[string]string{
		"id": "id2384569",
	})

	myData := make(map[string]interface{})
	myData["Name"] = "Tony"
	myData["Age"] = 23

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

	fmt.Println(map[string]interface{}{
		"title": "JWT 接口测试",
		"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiaWQiOiIxIiwibmFtZSI6IkhvYm8ifQ.YUzBykoELyKoQWaugkVNf3d09HBhICBJoOcWQKnveRQ",
	})

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

}

func privateFunc(a ...interface{}) {
	fmt.Println(a)
	fmt.Println(a...)
	for _, param := range a {
		fmt.Println(param)
	}
}
