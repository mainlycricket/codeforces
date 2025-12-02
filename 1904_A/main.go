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
		a, b := lineToNums(scanner.Text())

		scanner.Scan()
		xk, yk := lineToNums(scanner.Text())

		scanner.Scan()
		xq, yq := lineToNums(scanner.Text())

		var count int

		if killsQueen(xk+a, yk+b, a, b, xq, yq) {
			count++
		}
		if killsQueen(xk-a, yk+b, a, b, xq, yq) {
			count++
		}
		if killsQueen(xk-a, yk-b, a, b, xq, yq) {
			count++
		}
		if killsQueen(xk+a, yk-b, a, b, xq, yq) {
			count++
		}

		if a != b {
			if killsQueen(xk+b, yk+a, b, a, xq, yq) {
				count++
			}
			if killsQueen(xk-b, yk+a, b, a, xq, yq) {
				count++
			}
			if killsQueen(xk-b, yk-a, b, a, xq, yq) {
				count++
			}
			if killsQueen(xk+b, yk-a, b, a, xq, yq) {
				count++
			}
		}

		fmt.Println(count)
	}
}

func killsQueen(x, y, a, b, tx, ty int) bool {
	return ((x+a) == tx && (y+b) == ty) ||
		((x-a) == tx && (y+b) == ty) ||
		((x-a) == tx && (y-b) == ty) ||
		((x+a) == tx && (y-b) == ty) ||
		((x+b) == tx && (y+a) == ty) ||
		((x-b) == tx && (y+a) == ty) ||
		((x-b) == tx && (y-a) == ty) ||
		((x+b) == tx && (y-a) == ty)
}

func lineToNums(line string) (int, int) {
	strArr := strings.Split(line, " ")
	a, _ := strconv.Atoi(strArr[0])
	b, _ := strconv.Atoi(strArr[1])
	return a, b
}
