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

func getMoneyAmount(n int) int {
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for j := 1; j < n; j++ {
		for i := j - 1; i >= 0; i-- {
			if i+1 == j {
				res[i][j] = i + 1
			} else {
				initial := true
				for k := i + 1; k < j; k++ {
					localMax := k + 1 + getMax(res[i][k-1], res[k+1][j])
					if initial || localMax < res[i][j] {
						res[i][j] = localMax
						initial = false
					}
				}
			}
		}
	}
	return res[0][n-1]
}

func main() {
	fmt.Println(getMoneyAmount(62) == 222)
}
