package text

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Animation string

const Basic Animation = `-/|\`

type MessageTemplate int

const (
	Get MessageTemplate = iota
	Install
)

var scan = bufio.NewScanner(os.Stdin)

func Spinner(text string, spinStr Animation, delay time.Duration, done <-chan bool) {
	spinChars := []rune(spinStr)
	n := len(spinChars)
	idx := 0
	for {
		select {
		case <-done:
			return
		default:
			fmt.Printf("\r%s%c", text, spinChars[idx])
			idx = (idx + 1) % n
			time.Sleep(delay)
		}
	}
}

func ChoosePkg(pkgs []string, msgtype MessageTemplate) (string, bool) {
	n := len(pkgs)
	if n == 0 {
		panic("pkgs shouldn't be empty")
	}

	if n == 1 {
		pkg := pkgs[0]
		return pkg, getSingleAnswer(pkg, msgtype)
	}

	switch msgtype {
	case Get:
		fmt.Println("Choose a package to get:")
	case Install:
		fmt.Println("Choose a package to install:")
	default:
		fmt.Println("Choose a package:")
	}

	for i, p := range pkgs {
		fmt.Printf("%d. %s\n", i+1, p)
	}

	fmt.Println()

	num, ok := getNum(len(pkgs))
	if !ok {
		return "", false
	}

	return pkgs[num-1], true
}

func getSingleAnswer(pkg string, msgtype MessageTemplate) bool {
	switch msgtype {
	case Get:
		fmt.Printf("Get package %s? (y/n)\n> ", pkg)
	case Install:
		fmt.Printf("Install package %s? (y/n)\n> ", pkg)
	default:
		fmt.Printf("Choose package %s? (y/n)\n> ", pkg)
	}

	scan.Scan()
	ans := strings.ToLower(scan.Text())

	return ans == "y" || ans == "yes" || ans == ""
}

func getNum(opts int) (int, bool) {
	var ans string
	var err error
	var n int

	msg := fmt.Sprintf("(1..%d/n)\n> ", opts)

	fmt.Print(msg)

	for {
		scan.Scan()

		ans = strings.ToLower(scan.Text())

		if ans == "" {
			return 1, true
		}

		if ans == "n" || ans == "no" {
			return -1, false
		}

		n, err = strconv.Atoi(ans)
		if err != nil {
			fmt.Printf("Choose an option from %s", msg)
			continue
		}

		if n < 1 || n > opts {
			fmt.Printf("Number out of range, choose an option from %s", msg)
			continue
		}

		return n, true
	}
}

func ChooseInstallVersion(pkg string) (string, bool) {
	prompt := fmt.Sprintf(`Choose version of %s to install:
1. @latest
2. other
3. cancel

> `, pkg)

	fmt.Println()

	for {
		fmt.Print(prompt)

		scan.Scan()
		switch scan.Text() {
		case "", "1":
			return "@latest", true
		case "2":
			return getOtherVersion()
		case "3":
			return "", false
		}
	}
}

func getOtherVersion() (string, bool) {
	fmt.Print("Write the version you want (or /q to cancel):\n> ")
	scan.Scan()

	ans := scan.Text()

	if ans == "/q" || ans == "" {
		return "", false
	}

	if strings.HasPrefix(ans, "@") {
		return ans, true
	}

	return fmt.Sprintf("@%s", ans), true
}
