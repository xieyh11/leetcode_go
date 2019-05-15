package main

import (
	"fmt"
)

func lastRemaining(n int) int {
	if n < 1 {
		return 0
	}
	if n == 1 {
		return 1
	}
	earseType := []bool{} //false earse odd, true earse even
	direction := false    // false left to right, true right to left
	for n != 1 {
		if !direction || n&1 == 1 {
			earseType = append(earseType, false)
			n /= 2
		} else {
			earseType = append(earseType, true)
			n = (n + 1) / 2
		}
		direction = !direction
	}

	res := 1

	for i := len(earseType) - 1; i >= 0; i-- {
		if earseType[i] {
			res = res*2 - 1
		} else {
			res *= 2
		}
	}
	return res
}

func main() {
	fmt.Println(lastRemaining(9))
	fmt.Println(lastRemaining(11))
	fmt.Println(lastRemaining(100))
}
