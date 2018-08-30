package main

import (
	"fmt"
)

func multiply(num1 string, num2 string) string {
	num1L, num2L := len(num1), len(num2)
	res := make([]byte, num1L+num2L)
	for i := 0; i < num1L; i++ {
		for j := 0; j < num2L; j++ {
			num1B, num2B := byte(num1[num1L-1-i]-'0'), byte(num2[num2L-1-j]-'0')
			carry := num1B * num2B
			ijC := i + j
			for carry != 0 {
				carry = carry + res[ijC]
				res[ijC] = carry % 10
				ijC++
				carry /= 10
			}
		}
	}
	for i := range res {
		res[i] += '0'
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	i := 0
	for ; i < len(res)-1 && res[i] == '0'; i++ {
	}
	res = res[i:]
	return string(res)
}

func main() {
	num1 := "123"
	num2 := "456"
	fmt.Println(multiply(num1, num2))
}
