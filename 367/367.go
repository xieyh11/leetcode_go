package main

import (
	"fmt"
)

func isPerfectSquare(num int) bool {
	left := 1
	right := num

	for left < right-1 {
		mid := (left + right) / 2
		powerOf := mid * mid
		if powerOf == num {
			return true
		} else if powerOf < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left == right {
		return (left * left) == num
	}
	if left == right-1 {
		if left*left == num {
			return true
		}
		return right*right == num
	}
	return false
}

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i, isPerfectSquare(i))
	}
}
