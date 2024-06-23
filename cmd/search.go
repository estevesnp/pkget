package cmd

import (
	"fmt"

	"github.com/estevesnp/pkgo/internal/fetch"

	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [pkg]",
	Short: "Search for matching packages",
	Long: `Search for and list packages that match the provided package name. 
Choose the number of packages to display with the --limit or -l flag.`,
	Args:    cobra.ExactArgs(1),
	PreRunE: checkLimit,
	Run:     searchRun,
}

func init() {
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 5, "limit of packages displayed (1 <= limit <= 100)")
	rootCmd.AddCommand(searchCmd)
}

func searchRun(cmd *cobra.Command, args []string) {
	pkgArg := args[0]

	pkgs, err := fetch.SpinWhileFetching(pkgArg, limit)
	cobra.CheckErr(err)

	if len(pkgs) == 0 {
		fmt.Println("No packages found")
		return
	}

	for _, p := range pkgs {
		fmt.Println(p)
	}
}
