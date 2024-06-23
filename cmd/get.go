package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/estevesnp/pkgo/internal/fetch"
	"github.com/estevesnp/pkgo/internal/text"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [pkg]",
	Short: "Add package to project",
	Long: `Add a package to your Go project, similarly to doing "go get [pkg]". Requires having Go installed.
You can define a limit of displayed packages with the -l flag, pass a version with the -v flag and update the package dependencies with the -u flag.`,
	Args:    cobra.ExactArgs(1),
	PreRunE: checkAll(checkPath, checkLimit),
	Run:     getRun,
}

func init() {
	getCmd.Flags().IntVarP(&limit, "limit", "l", 5, "limit of packages displayed (1 <= limit <= 100)")
	getCmd.Flags().StringVarP(&version, "version", "v", "", "define a version for the package")
	getCmd.Flags().BoolVarP(&update, "update", "u", false, "update pkg dependencies")
	rootCmd.AddCommand(getCmd)
}

func getRun(cmd *cobra.Command, args []string) {
	pkgArg := args[0]

	pkgs, err := fetch.SpinWhileFetching(pkgArg, limit)
	cobra.CheckErr(err)

	n := len(pkgs)
	if n == 0 {
		fmt.Println("No packages found")
		return
	}

	pkg, ok := text.ChoosePkg(pkgs, text.Get)
	if !ok {
		fmt.Println("No package selected")
		return
	}

	version, _ = verifyVersion(version)

	cobra.CheckErr(goGet(pkg, version, update))
}

func goGet(pkg, version string, update bool) error {
	fullPkg := fmt.Sprintf("%s%s", pkg, version)

	args := []string{"get"}
	if update {
		args = append(args, "-u")
	}
	args = append(args, fullPkg)

	fmt.Printf("\nRunning go %s...\n", strings.Join(args, " "))

	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
