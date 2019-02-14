package main

import (
	"fmt"
)

func generate(numRows int) [][]int {
	if numRows <= 0 {
		return [][]int{}
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][len(res[i])-1] = 1
		for j := 1; j < len(res[i])-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}

func main() {
	fmt.Println(generate(5))
}
