package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testcases, _ := strconv.Atoi(scanner.Text())

	for range testcases {
		// read and discard 'n'
		scanner.Scan()
		scanner.Text()

		scanner.Scan()
		res := fn(scanner.Text())
		fmt.Println(res)
	}
}

func fn(input string) int {
	if strings.Contains(input, "...") {
		return 2
	}
	return strings.Count(input, ".")
}
