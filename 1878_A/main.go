package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	testcases, _ := strconv.Atoi(scanner.Text())

	for range testcases {
		scanner.Scan()
		info := lineToNums(scanner.Text())

		scanner.Scan()
		nums := lineToNums(scanner.Text())

		res := fn(nums, info[1])
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func fn(nums []int, k int) bool {
	return slices.Contains(nums, k)
}

func lineToNums(line string) []int {
	strArr := strings.Split(line, " ")
	intArr := make([]int, len(strArr))
	for idx, value := range strArr {
		intArr[idx], _ = strconv.Atoi(value)
	}
	return intArr
}
