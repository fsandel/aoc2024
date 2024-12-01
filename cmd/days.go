package cmd

import (
	"aoc/internal/day1"

	"github.com/spf13/cobra"
)

var day1Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1 puzzle",
	Long:  "Calculate difference and similarity of 2 lists",
	RunE: func(cmd *cobra.Command, args []string) error {
		return day1.Run()
	},
}

func init() {
	rootCmd.AddCommand(day1Cmd)
}
