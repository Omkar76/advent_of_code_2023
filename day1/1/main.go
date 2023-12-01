package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func parse(s string) int {
	first := 0
	for i := 0; i < len(s); i++ {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			first = int(c - '0')
			break
		}
	}

	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if unicode.IsDigit(c) {
			last = int(c - '0')
			break
		}
	}

	return first*10 + last

}
func main() {

	r := bufio.NewScanner(os.Stdin)
	sum := 0
	for r.Scan() {
		sum += parse(r.Text())
	}

	fmt.Println(sum)
}
