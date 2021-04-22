package cmd

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	worldCmd.AddCommand(worldAddCmd)
	worldCmd.AddCommand(worldPrintCmd)
}

var worldCmd = &cobra.Command{
	Use:   "world",
	Short: "Tools for managing a world",
	Run: func(cmd *cobra.Command, args []string) {
		world()
	},
}

func world() {
	fmt.Println("hello, world")
}

var worldAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Create the skeleton of a new world.",
	Run: func(cmd *cobra.Command, args []string) {
		worldAdd(args[0])
	},
}

func worldAdd(name string) {
	var err error

	if _, err = os.Stat(fmt.Sprintf("world/%s.yaml", name)); os.IsNotExist(err) {
		file, err := os.Create(fmt.Sprintf("world/%s.yaml", name))
		if err != nil {
			log.Fatal(err)
		}

		current := World{}
		current.Name = name

		data, err := yaml.Marshal(&current)
		if err != nil {
			log.Fatal(err)
		}

		_, err = io.WriteString(file, string(data))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Printf("world %s exists", name)
	}
}

var worldPrintCmd = &cobra.Command{
	Use:   "print <name>",
	Short: "Print overview of world.",
	Run: func(cmd *cobra.Command, args []string) {
		worldPrint(args[0])
	},
}

func worldPrint(name string) {
	var err error

	if _, err = os.Stat(fmt.Sprintf("world/%s.yaml", name)); err == nil {
		current := World{}
		current.Name = name

		data, err := ioutil.ReadFile(fmt.Sprintf("world/%s.yaml", name))
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(data, &current)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%+v", current)
	} else {
		log.Fatalf("world %s does exist", name)
	}
}

// World is a structure holding yaml properties for world
type World struct {
	Name string
}
