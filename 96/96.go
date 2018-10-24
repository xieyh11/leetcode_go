package main

import (
	"fmt"
)

func numTrees(n int) int {
	res := make([]int, n+1)
	res[0] = 1
	res[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			res[i] += (res[j] * res[i-1-j])
		}
	}
	return res[n]
}

func main() {
	fmt.Println(numTrees(3))
}
