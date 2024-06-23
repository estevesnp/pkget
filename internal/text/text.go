package text

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const Basic = `-/|\`

func Spinner(text, spinStr string, delay time.Duration, done <-chan bool) {
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

func ChoosePkg(pkgs []string) (string, bool) {
	n := len(pkgs)
	if n == 0 {
		panic("pkgs shouldn't be empty")
	}

	if n == 1 {
		pkg := pkgs[0]
		return pkg, getSingleAnswer(pkg)
	}

	fmt.Println("Choose a package:")

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

func getSingleAnswer(pkg string) bool {
	s := bufio.NewScanner(os.Stdin)
	fmt.Printf("Use package %s? (y/n)\n> ", pkg)
	s.Scan()
	ans := strings.ToLower(s.Text())

	return ans == "y" || ans == "yes" || ans == ""
}

func getNum(opts int) (int, bool) {
	var ans string
	var err error
	var n int

	s := bufio.NewScanner(os.Stdin)

	msg := fmt.Sprintf("(1..%d/n)\n> ", opts)

	fmt.Print(msg)

	for {
		s.Scan()

		ans = s.Text()

		if ans == "" {
			return 1, true
		}

		if ans == "n" || ans == "N" {
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
