package main

import (
	"fmt"
)

func recSol(k int, n int, min int) [][]int {
	if k == 1 {
		if n >= min && n <= 9 {
			return [][]int{{n}}
		} else {
			return [][]int{}
		}
	}
	res := make([][]int, 0)
	for i := min; i <= 9-k+1; i++ {
		temp := recSol(k-1, n-i, i+1)
		for _, one := range temp {
			res = append(res, append([]int{i}, one...))
		}
	}
	return res
}

func combinationSum3(k int, n int) [][]int {
	if k == 0 || k > 9 {
		return [][]int{}
	}
	return recSol(k, n, 1)
}

func main() {
	fmt.Println(combinationSum3(7, 35))
}
