package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	arr := lineToNums(scanner.Text())
	n, m, a, b := arr[0], arr[1], arr[2], arr[3]
	var res int
	if (m / b) < a {
		y := int(n) / int(m)
		res += y * int(b)
		n -= m * float64(y)
	}
	res += min(int(n)*int(a), int(math.Ceil((n/m))*b))
	fmt.Println(res)
}

func lineToNums(line string) []float64 {
	strArr := strings.Split(line, " ")
	numArr := make([]float64, len(strArr))
	for idx, value := range strArr {
		numArr[idx], _ = strconv.ParseFloat(value, 64)
	}
	return numArr
}
