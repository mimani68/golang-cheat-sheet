package main

import "fmt"

type SharedState struct {
	Liveness bool
}

type Database struct {
	Address string
}

// Note: Go does not have traditional classes or inheritance.
// Instead, it uses composition By embedding
type Application struct {
	Title        string // Named Field, It is not embedded
	privateTitle string
	Database     // Struct Embedding (Composition)
	*SharedState // Composition design pattern
}

func main() {

	d := Database{
		Address: "tcp://sc33ff489vh3.mongodb.com/1",
	}

	state := SharedState{
		Liveness: true,
	}

	app := Application{
		privateTitle: "simple application",
		Database:     d,
		SharedState:  &state,
	}

	fmt.Println(app)
	fmt.Printf("%v\n", app.SharedState)
	fmt.Printf("%v\n", app.privateTitle)

}

// Output
// { {tcp://sc33ff489vh3.mongodb.com/1} 0xc0000120fd}
// &{true}
