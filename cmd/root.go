package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkget",
	Short: "Find and install Go packages",
	Long:  "pkget is a simple CLI tool to help you find, install and manage Go packages.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
		_ = cmd.Usage()
		os.Exit(1)
	},
}

func Execute() {
	_ = rootCmd.Execute()
}
