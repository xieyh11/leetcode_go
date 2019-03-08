package main

import (
	"fmt"
)

var numPow = []int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}

func next(n int) int {
	res := 0
	for n != 0 {
		res += numPow[n%10]
		n /= 10
	}
	return res
}

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}
	one := n
	two := n
	for two != 1 {
		one = next(one)
		two = next(next(two))
		if one == two {
			return two == 1
		}
	}
	return true
}

func main() {
	fmt.Println(isHappy(19))
}
