package main

import (
	"fmt"
)

func rotate(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	row := len(matrix)
	for i := 0; i < row; i++ {
		for j := i; j < row; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	half := row / 2
	for r := 0; r < row; r++ {
		for c := 0; c < half; c++ {
			matrix[r][c], matrix[r][row-1-c] = matrix[r][row-1-c], matrix[r][c]
		}
	}
}

func main() {
	matrix := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(matrix)
	rotate(matrix)
	fmt.Println(matrix)
}
