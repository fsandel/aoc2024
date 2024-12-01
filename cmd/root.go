package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "Root",
	Long:  "Root CMD",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to my implementation of of the advent of code 2024")
		fmt.Println("To get the solutions of the days, use them as commands")
		fmt.Println("e.g. go run main.go day1")
	},
}

func Execute() error {
	return rootCmd.Execute()
}
