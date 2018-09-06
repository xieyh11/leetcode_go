package main

import (
	"fmt"
)

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	a0 := 1
	a1 := 2
	a2 := 0
	for i := 2; i < n; i++ {
		a2 = a0 + a1
		a0, a1 = a1, a2
	}
	return a2
}

func main() {
	fmt.Println(climbStairs(4))
}
