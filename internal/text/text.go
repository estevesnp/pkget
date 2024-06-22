package text

import (
	"fmt"
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
