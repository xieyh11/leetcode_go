package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	carry := 1
	for i := len(digits) - 1; i >= 0 && carry != 0; i-- {
		digits[i] += carry
		carry = digits[i] / 10
		digits[i] %= 10
	}
	if carry != 0 {
		digits = append([]int{carry}, digits...)
	}
	return digits
}

func main() {
	digits := []int{1, 2, 3}
	fmt.Println(plusOne(digits))
}
