package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

var words = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
}

func parse(s string) int {
	n := len(s)
	first := 0

first_out:
	for i := 0; i < n; i++ {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			first = int(c - '0')
			break
		} else {
			for d, word := range words {
				if strings.HasPrefix(s[i:n], word) {
					first = d
					break first_out
				}
			}
		}

	}

	last := 0
last_out:
	for i := n - 1; i >= 0; i-- {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			last = int(c - '0')
			break
		} else {
			for d, word := range words {
				if strings.HasPrefix(s[i:n], word) {
					last = d
					break last_out
				}
			}
		}

	}
	return first*10 + last

}
func main() {

	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan() {
		s := scanner.Text()
		n := parse(s)
		sum += n
	}

	fmt.Println(sum)
}
