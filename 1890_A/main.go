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
		if res {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func fn(nums []int) bool {
	var first, second int
	var firstCount, secondCount int

	for _, num := range nums {
		if num == first {
			firstCount++
		} else if num == second {
			secondCount++
		} else if first == 0 {
			first = num
			firstCount++
		} else if second == 0 {
			second = num
			secondCount++
		} else {
			return false
		}
	}

	if secondCount == 0 {
		return true
	}

	diff := firstCount - secondCount
	return diff >= -1 && diff <= 1
}

func lineToNums(line string) []int {
	strArr := strings.Split(line, " ")
	intArr := make([]int, len(strArr))
	for idx, value := range strArr {
		intArr[idx], _ = strconv.Atoi(value)
	}
	return intArr
}
