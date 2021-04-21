// Package main provides ...
package main

import (
	"fmt"

	"github.com/alecthomas/kong"
)

// CLI is the struct containing all arguments and help
var CLI struct {
	Add struct {
	} `cmd help:"Add an item"`
	Rm struct {
	} `cmd help:"Remove an item"`
	Edit struct {
	} `cmd help:"Edit and item"`
	Rs struct {
	} `cmd help:"List items"`
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
