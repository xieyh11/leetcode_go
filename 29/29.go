package main

import (
	"fmt"
)

const MaxInt = 1<<31 - 1

func divide(dividend int, divisor int) int {
	count := 0
	if dividend >= 0 && divisor > 0 {
		for dividend >= divisor {
			dividend -= divisor
			count++
		}
	} else if dividend < 0 || (dividend >= 0 && divisor < 0) {
		positive := true
		if dividend < 0 && divisor > 0 {
			positive = false
			divisor = -divisor
		} else if dividend >= 0 && divisor < 0 {
			positive = false
			dividend = -dividend
		}
		for dividend <= divisor {
			dividend -= divisor
			count++
		}
		if !positive {
			count = -count
		}
	}
	if count > MaxInt {
		count = MaxInt
	}
	return count
}

func main() {
	fmt.Println(divide(-2147483648, 1))
}
