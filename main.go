// Package main provides ...
package main

import (
	"fmt"
	"log"

	"github.com/alecthomas/kong"
	"gopkg.in/yaml.v2"
)

// World is a struct containing the variables and external structs of a world.
type World struct {
	Name string
}

// CLI is the struct containing all arguments and help
var CLI struct {
	Add struct {
	} `cmd help:"Add an item"`
	Rm struct {
	} `cmd help:"Remove an item"`
}

func oldmain() {
	world := World{}
	world.Name = "foo"

	doc, err := yaml.Marshal(&world)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(doc))

	err = yaml.Unmarshal([]byte(doc), &world)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", world)

	return
}

func main() {
	ctx := kong.Parse(&CLI)
	switch ctx.Command() {
	case "add":
		fmt.Println("Foo")
	default:
		panic(ctx.Command())
	}
}
