package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/estevesnp/pkget/internal/fetch"
	"github.com/estevesnp/pkget/internal/text"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [pkg]",
	Short: "Add package to project",
	Long: `Add a package to your Go project, similarly to doing "go get [pkg]". Requires having Go installed.
You can pass some of the same flags as you would pass the "go get" command.`,
	Args:    cobra.ExactArgs(1),
	PreRunE: checkAll(checkPath, checkLimit),
	RunE:    getRun,
}

func init() {
	getCmd.Flags().BoolVarP(&update, "update", "u", false, "update pkg dependencies")
	getCmd.Flags().IntVarP(&limit, "limit", "l", 5, "limit of packages displayed (1 <= limit <= 100)")
	rootCmd.AddCommand(getCmd)
}

func getRun(cmd *cobra.Command, args []string) error {
	pkgArg := args[0]

	pkgs, err := fetch.SpinWhileFetching(pkgArg, limit)
	cobra.CheckErr(err)

	n := len(pkgs)
	if n == 0 {
		fmt.Println("No packages found")
		return nil
	}

	pkg, ok := text.ChoosePkg(pkgs, text.Get)
	if !ok {
		return nil
	}

	return goGet(pkg, update)
}

func goGet(pkg string, update bool) error {
	args := []string{"get"}
	if update {
		args = append(args, "-u")
	}
	args = append(args, pkg)

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
