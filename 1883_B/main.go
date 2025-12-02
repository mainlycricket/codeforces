package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	testcases, _ := strconv.Atoi(readLine())

	for range testcases {
		n, k := lineToNums(readLine())
		res := fn(readLine(), n, k)
		if res {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func fn(input string, n, k int) bool {
	var (
		count       [26]int
		oddCount    int
		isFinalEven = (n-k)%2 == 0
	)

	for _, ch := range input {
		idx := ch - 'a'
		count[idx]++
	}

	for idx, chCount := range count {
		if k > 0 && chCount%2 == 1 {
			count[idx]--
			k--
		}
		if count[idx]%2 == 1 {
			oddCount++
		}
	}

	if isFinalEven {
		return oddCount == 0 && k%2 == 0
	}

	for idx, chCount := range count {
		if chCount%2 == 0 && k > 1 {
			diff := min(chCount, k)
			if diff%2 == 1 {
				diff--
			}
			count[idx] -= diff
			k -= diff
		}
	}

	return (oddCount == 1 && k == 0) || (oddCount == 0 && k == 1)
}

func lineToNums(line string) (int, int) {
	strArr := strings.Split(line, " ")
	a, _ := strconv.Atoi(strArr[0])
	b, _ := strconv.Atoi(strArr[1])
	return a, b
}

func readLine() string {
	size := 100000
	arr := make([]byte, size)

	reader := bufio.NewReaderSize(os.Stdin, size)
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(arr, size)
	scanner.Scan()

	line := bytes.Split(arr, []byte("\n"))
	text := string(line[0])
	return text
}
