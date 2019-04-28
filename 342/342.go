package main

import (
	"fmt"
)

func isPowerOfFour(num int) bool {
	for num > 0 {
		if num == 1 {
			return true
		}
		if num&1 != 0 {
			return false
		}
		if num&2 != 0 {
			return false
		}
		num >>= 2
	}
	return false
}

func main() {
	for i := 1; i < 50; i++ {
		fmt.Println(i, isPowerOfFour(i))
	}
}
