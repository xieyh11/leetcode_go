package main

import (
	"fmt"
)

var Res = []int{10, 91, 739, 5275, 32491, 168571, 712891, 2345851, 5611771, 8877691}

func countNumbersWithUniqueDigits(n int) int {
	if n > 10 {
		return Res[9]
	}
	if n == 0 {
		return 1
	}
	return Res[n-1]
}

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
}
