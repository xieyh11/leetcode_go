package main

import (
	"fmt"
)

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	n &= n - 1
	return n == 0
}

func main() {
	fmt.Println(isPowerOfTwo(16))
}
