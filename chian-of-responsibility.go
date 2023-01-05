package main

import "fmt"

type SDK struct {
	data string
}

func (app *SDK) leftJoin() *SDK {
	fmt.Println("leftJoin")
	return app
}

func (app *SDK) mapData() *SDK {
	fmt.Println("Map")
	app.data = "sample data"
	return app
}

func (app *SDK) returnValur() string {
	fmt.Println("Return")
	return app.data
}

func main() {
	var app = SDK{}
	app.leftJoin()
	app.mapData()
	app.leftJoin()
	fmt.Println(app.returnValur())
}
