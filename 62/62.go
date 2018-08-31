package main

import (
	"fmt"
)

func nChooseK(n, k int) int {
	if n < k {
		return 0
	}
	if n-k < k {
		k = n - k
	}
	if k == 0 {
		return 1
	}
	upFac := 1
	doFac := 1
	for i, j := 1, n-k+1; i <= k; i, j = i+1, j+1 {
		upFac *= j
		doFac *= i
	}
	return upFac / doFac
}

func uniquePaths(m int, n int) int {
	return nChooseK(m-1+n-1, n-1)
}

func main() {
	m := 3
	n := 2
	fmt.Println(uniquePaths(m, n))
}
