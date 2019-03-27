package main

import (
	"fmt"
)

func remainFactor(num, factor int) int {
	for num%factor == 0 {
		num /= factor
	}
	return num
}

func isUgly(num int) bool {
	num = remainFactor(num, 2)
	num = remainFactor(num, 3)
	num = remainFactor(num, 5)
	return num == 1
}

func main() {
	fmt.Println(isUgly(14))
}
