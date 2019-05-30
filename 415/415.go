package main

import (
	"fmt"
)

func getCarry(num byte) (byte, byte) {
	if num > '9' {
		return num - 10, 1
	} else {
		return num, 0
	}
}

func addStrings(num1 string, num2 string) string {
	maxLen := len(num1)
	if maxLen < len(num2) {
		maxLen = len(num2)
	}
	res := make([]byte, maxLen)
	carry := byte(0)
	last1 := len(num1) - 1
	last2 := len(num2) - 1
	lastR := len(res) - 1
	for last1 >= 0 && last2 >= 0 {
		sum := num1[last1] - '0' + num2[last2] + carry
		sum, carry = getCarry(sum)
		res[lastR] = sum
		last1--
		last2--
		lastR--
	}
	for last1 >= 0 {
		sum := num1[last1] + carry
		sum, carry = getCarry(sum)
		res[lastR] = sum
		last1--
		lastR--
	}
	for last2 >= 0 {
		sum := num2[last2] + carry
		sum, carry = getCarry(sum)
		res[lastR] = sum
		last2--
		lastR--
	}
	if carry != 0 {
		return string(carry+'0') + string(res)
	} else {
		return string(res)
	}
}

func main() {
	num1 := "122"
	num2 := "243"
	fmt.Println(addStrings(num1, num2))

	num1 = "99"
	num2 = "99"
	fmt.Println(addStrings(num1, num2))
}
