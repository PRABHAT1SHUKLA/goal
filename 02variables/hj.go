package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{Use: "mycli"}

	var greetCmd = &cobra.Command{
		Use:   "greet",
		Short: "Prints a greeting",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, welcome to our CLI!")
		},
	}

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Displays the CLI version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("CLI Version 1.0.0")
		},
	}

	rootCmd.AddCommand(greetCmd, versionCmd)
	rootCmd.Execute()
}
