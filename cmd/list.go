package cmd

import (
	"fmt"
	"time"

	"github.com/estevesnp/pkget/internal/fetch"
	"github.com/estevesnp/pkget/internal/text"

	"github.com/spf13/cobra"
)

var limit int

var listCmd = &cobra.Command{
	Use:   "list [pkg]",
	Short: "List matching packages",
	Long: `List packages that match the provided package name. 
Choose the number of packages to display with the --limit or -l flag.`,
	Args:    cobra.ExactArgs(1),
	PreRunE: listValidate,
	Run:     listRun,
}

func init() {
	listCmd.Flags().IntVarP(&limit, "limit", "l", 5, "limit of packages displayed (1 <= limit <= 100)")
	rootCmd.AddCommand(listCmd)
}

func listValidate(cmd *cobra.Command, args []string) error {
	if limit < 1 || limit > 100 {
		return fmt.Errorf("invalid limit: %d (must be between 1 and 100)", limit)
	}
	return nil
}

func listRun(cmd *cobra.Command, args []string) {
	pkg := args[0]

	done := make(chan bool)

	txt := fmt.Sprintf("Fetching packages for %q ", pkg)
	go text.Spinner(txt, text.Basic, 100*time.Millisecond, done)

	pkgs, found, err := fetch.FetchPackages(pkg, limit)
	done <- true
	fmt.Print("\n\n")
	cobra.CheckErr(err)

	if !found {
		fmt.Println("No packages found")
		return
	}

	for _, p := range pkgs {
		fmt.Println(p)
	}
}
