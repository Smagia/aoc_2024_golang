package cmd

import (
	"aoc/smagia/internal/day6"
	"aoc/smagia/internal/day7"
	"aoc/smagia/internal/day8"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "run",
	Short: "run required exercise",
	Long:  "run the specified exercise",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("You asked for: %s\n", args[0])
		switch args[0] {
		case "day6":
			day6.Run1()
			day6.Run2()
		case "day7":
			day7.Run1()
			day7.Run2()
		case "day8":
			day8.Run1()
			day8.Run2()
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("Error while running command '%s'", err)
		os.Exit(1)
	}
}
