package cmd

import (
	"fmt"

	"github.com/estevesnp/pkget/internal/fetch"
	"github.com/spf13/cobra"
)

var limit int

var listCmd = &cobra.Command{
	Use:   "list [pkg]",
	Short: "List matching packages",
	Long: `List packages that match the provided package name.
For example:

pkgo list -l 10 gin`,
	Args: cobra.ExactArgs(1),
	Run:  list,
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&limit, "limit", "l", 5, "limit of packages displayed")
}

func list(cmd *cobra.Command, args []string) {
	pkg := args[0]

	fmt.Printf("Fetching packages for %q...\n", pkg)

	pkgs, found, err := fetch.FetchPackages(pkg, limit)
	cobra.CheckErr(err)

	if !found {
		fmt.Println("No packages found")
		return
	}

	for _, p := range pkgs {
		fmt.Println(p)
	}
}
