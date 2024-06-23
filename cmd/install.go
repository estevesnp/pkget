package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/estevesnp/pkgo/internal/fetch"
	"github.com/estevesnp/pkgo/internal/text"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install [pkg]",
	Short: "Install a package",
	Long: `Install a package to your Go Path, similarly to doing "install get [pkg]". Requires having Go installed.
You can define a limit of displayed packages with the -l flag and pass a version with the -v flag.`,
	Args:    cobra.ExactArgs(1),
	PreRunE: checkAll(checkPath, checkLimit),
	Run:     installRun,
}

func init() {
	installCmd.Flags().IntVarP(&limit, "limit", "l", 5, "limit of packages displayed (1 <= limit <= 100)")
	installCmd.Flags().StringVarP(&version, "version", "v", "", "define a version for the package")
	rootCmd.AddCommand(installCmd)
}

func installRun(cmd *cobra.Command, args []string) {
	pkgArg := args[0]

	pkgs, err := fetch.SpinWhileFetching(pkgArg, limit)
	cobra.CheckErr(err)

	n := len(pkgs)
	if n == 0 {
		fmt.Println("No packages found")
		return
	}

	pkg, ok := text.ChoosePkg(pkgs, text.Install)
	if !ok {
		fmt.Println("No package selected")
		return
	}

	version, ok = verifyVersion(version)
	if !ok {
		version, ok = text.ChooseInstallVersion(pkg)
		if !ok {
			fmt.Println("No version selected")
			return
		}
	}

	cobra.CheckErr(goInstall(pkg, version))
}

func goInstall(pkg, version string) error {
	fullPkg := fmt.Sprintf("%s%s", pkg, version)

	fmt.Printf("\nRunning go install %s...\n", fullPkg)

	cmd := exec.Command("go", "install", fullPkg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
