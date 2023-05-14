package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	tmpl, err := template.New("test").Parse(`Hello {{.Username}}!`)
	if err != nil {
		fmt.Errorf("&s", err.Error())
		return
	}

	// => Correct
	// data := struct {
	// 	Username string
	// }{
	// 	Username: "John Doe",
	// }

	// => Correct
	data := map[string]string{
		"Username": "John Doe",
	}

	if err := tmpl.Execute(os.Stdout, data); err != nil {
		fmt.Errorf("&s", err.Error())
		return
	}
}
