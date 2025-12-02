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
		nums := lineToNums(scanner.Text())

		res := fn(nums)
		fmt.Println(res)
	}
}

func fn(nums []int) int {
	var sum int

	for _, num := range nums {
		sum += num
	}

	return sum * -1
}

func lineToNums(line string) []int {
	strArr := strings.Split(line, " ")
	intArr := make([]int, len(strArr))
	for idx, value := range strArr {
		intArr[idx], _ = strconv.Atoi(value)
	}
	return intArr
}
