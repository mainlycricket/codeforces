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
		scanner.Scan()
		info := lineToNums(scanner.Text())
		_, k := info[0], info[1]

		scanner.Scan()
		nums := lineToNums(scanner.Text())

		res := fn(nums, k)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func fn(nums []int, k int) bool {
	if k > 1 {
		return true
	}

	prev := nums[0]
	for _, num := range nums {
		if num < prev {
			return false
		}
		prev = num
	}
	return true
}

func lineToNums(line string) []int {
	strArr := strings.Split(line, " ")
	intArr := make([]int, len(strArr))
	for idx, value := range strArr {
		intArr[idx], _ = strconv.Atoi(value)
	}
	return intArr
}
