package main

import "fmt"

type Document struct {
	Title   string
	Pages   int
	Authors []string
}

func main() {
	var doc []Document
	doc = append(doc, Document{
		Title: "Mobidoc",
		Pages: 12,
		Authors: []string{
			"Shekspir",
		},
	})
	fmt.Println(doc[0].Authors)
	fmt.Println(doc)
}
