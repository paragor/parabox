package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(check([]rune(scanner.Text())))
}

func check(s []rune) int {
	curIsUpper := false
	sum := 0
	l := len(s)

	for i, x := range s {
		xIsUpper := isUpper(x)
		if x == ' ' {
			curIsUpper = false
			sum += 1
			continue
		}
		if xIsUpper == curIsUpper {
			sum += 1
			continue
		}
		if i+1 >= l {
			sum += 2
			continue
		}

		if s[i+1] == ' ' {
			sum += 2
			continue
		}

		if isUpper(s[i+1]) == xIsUpper {
			sum += 2
			curIsUpper = !curIsUpper
			continue
		}
		sum += 2
	}
	return sum
}
func isUpper(x rune) bool {
	return unicode.IsUpper(x)
}
