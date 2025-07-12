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
		n, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")

		nums := make([]int, len(arr))
		for idx, num := range arr {
			nums[idx], _ = strconv.Atoi(num)
		}

		res := fn(nums, n)
		fmt.Println(res)
	}
}

func fn(nums []int, n int) int {
	slices.Sort(nums)
	start, end := 0, len(nums)-1

	isStartEven := nums[start]%2 == 0
	isEndEven := nums[end]%2 == 0

	if isStartEven == isEndEven {
		return 0
	}

	for idx, num := range nums {
		if isStartEven && num%2 == 1 {
			start = idx
			break
		} else if !isStartEven && num%2 == 0 {
			start = idx
			break
		}
	}

	for idx := n - 1; idx >= 0; idx-- {
		if isEndEven && nums[idx]%2 == 1 {
			end = idx
			break
		} else if !isEndEven && nums[idx]%2 == 0 {
			end = idx
			break
		}
	}

	return min(start-0, (n-1)-end)
}
