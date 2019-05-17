package main

import (
	"fmt"
	"strconv"
)

func findNthDigit(n int) int {
	if n < 10 {
		return n
	}

	d, i := 9, 1
	startNum := 1
	for n > d*i {
		n -= d * i
		d *= 10
		i++
		startNum *= 10
	}
	num := (n-1)/i + startNum
	mod := (n - 1) % i
	str := strconv.FormatInt(int64(num), 10)
	return int(str[mod] - '0')
}

func main() {
	fmt.Println(findNthDigit(13))
}
