package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use: "init <path>",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			return nil
		}
		return fmt.Errorf("invalid path")
	},
	Short: "Initialize a rmanage repository.",
	Long:  "Initializes the skeleton of an rmanage repository.",
	Run: func(cmd *cobra.Command, args []string) {
		initRun(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initRun(path string) {

	var err error

	// Create directory
	os.Mkdir(path, os.ModePerm)

	// Create subdirectories through foreach
	directories := []string{"campaign", "character", "world"}
	for _, directory := range directories {
		if _, err = os.Stat(directory); err == nil {
			err = os.Mkdir(fmt.Sprintf("%s/%s", path, directory), os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Printf("%s exists", directory)
		}
	}

	// Create files
	files := []string{"config.yaml"}
	for _, file := range files {
		if _, err = os.Stat(file); err == nil {
			_, err = os.Create(fmt.Sprintf("%s/%s", path, file))
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Printf("%s exists", file)
		}
	}
}
