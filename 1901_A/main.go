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
		n, x := info[0], info[1]

		scanner.Scan()
		nums := lineToNums(scanner.Text())

		res := fn(nums, n, x)
		fmt.Println(res)
	}
}

func fn(nums []int, n, x int) int {
	res := (x - nums[n-1]) * 2
	cur := 0

	for _, num := range nums {
		res = max(res, num-cur)
		cur = num
	}

	return res
}

func lineToNums(line string) []int {
	strArr := strings.Split(line, " ")
	intArr := make([]int, len(strArr))
	for idx, value := range strArr {
		intArr[idx], _ = strconv.Atoi(value)
	}
	return intArr
}
