package cmd

import (
	// "fmt"

	// homedir "github.com/mitchellh/go-homedir"

	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string

	rootCmd = &cobra.Command{
		Use:   "rmanage",
		Short: "A tool to manage TTRPGs.",
		Long:  `A tool to manage TTPRGs.`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(worldCmd)

	// Set Config File
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("yaml")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")      // optionally look for config in the working directory

	if _, err := os.Stat("./config.yaml"); err == nil {
		err := viper.ReadInConfig() // Find and read the config file
		if err != nil {             // Handle errors reading the config file
			log.Fatal(err)
		}
	}
}
