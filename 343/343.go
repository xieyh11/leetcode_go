package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func integerBreak(n int) int {
	res := make([]int, n)
	res[0] = 0
	for i := 1; i < n; i++ {
		res[i] = getMax(res[i-1], (i+1)/2*(i+1-(i+1)/2))
		for j := 1; j <= i; j++ {
			res[i] = getMax(res[i], res[i-j]*j)
		}
	}
	return res[n-1]
}

func main() {
	fmt.Println(integerBreak(2))
	fmt.Println(integerBreak(10))
}
