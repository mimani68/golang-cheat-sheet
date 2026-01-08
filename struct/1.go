package main

import "fmt"

type SharedState struct {
	Liveness bool
}

type Database struct {
	Address string
}

type Config struct {
	Port int
}

type pullOperation struct {
	progress int
}

// Note: Go does not have traditional classes or inheritance.
// Instead, it uses composition By embedding
type Application struct {
	Title                    string // Named Field, It is not embedded
	privateTitle             string
	Database                 // Struct Embedding (Composition)
	*SharedState             // Composition design pattern
	Config                   *Config
	monitorsChan             chan struct{}
	Uid, Pid                 int64
	pullOperationsInProgress map[string]*pullOperation
}

func main() {

	config := Config{
		Port: 3000,
	}

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
		Config:       &config,
	}

	fmt.Println(app)
	fmt.Printf("%v\n", app.privateTitle)
	fmt.Printf("%v\n", app.SharedState)
	fmt.Printf("%v\n", app.Config)

}

// Output
// ==============
// { simple application {tcp://sc33ff489vh3.mongodb.com/1} 0xc000012108 0xc000012100}
// simple application
// &{true}
// &{3000}
