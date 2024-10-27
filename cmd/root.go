package cmd

import (
	"comet/scanner"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "comet",
	Short: "Comet is a CLI tool to scan and document comments in code files",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		fmt.Printf("Starting comment scan in directory: %s\n", path)
		scanner.ScanProject(path)

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
