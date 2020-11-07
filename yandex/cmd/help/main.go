package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var count int

	fmt.Scan(&count)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Text()

	for i := 0; i < count && scanner.Scan(); i++ {
		scanner.Text()
		scanner.Bytes()
	}
}
