package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	limit   int    // get, install, search
	version string // get, install
	update  bool   // get
)

var rootCmd = &cobra.Command{
	Use:   "pkgo",
	Short: "Find and install Go packages",
	Long:  "pkgo is a simple CLI tool to help you find, install and manage Go packages.",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Usage()
		os.Exit(1)
	},
}

func Execute() {
	_ = rootCmd.Execute()
}

func verifyVersion(version string) (string, bool) {
	if version == "" {
		return version, false
	}

	if !strings.HasPrefix(version, "@") {
		version = fmt.Sprintf("@%s", version)
	}

	return version, true
}

func checkLimit(_ *cobra.Command, _ []string) error {
	if limit < 1 || limit > 100 {
		return fmt.Errorf("invalid limit: %d (must be between 1 and 100)", limit)
	}
	return nil
}

func checkPath(_ *cobra.Command, _ []string) error {
	_, err := exec.LookPath("go")
	return err
}

func checkAll(funcs ...func(*cobra.Command, []string) error) func(*cobra.Command, []string) error {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range funcs {
			if err := f(cmd, args); err != nil {
				return err
			}
		}

		return nil
	}
}
