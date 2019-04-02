package main

import (
	"fmt"
)

type NumMatrix struct {
	Matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) > 0 && len(matrix[0]) > 0 {
		rows, cols := len(matrix), len(matrix[0])
		for i := 0; i < rows; i++ {
			for j := 1; j < cols; j++ {
				matrix[i][j] += matrix[i][j-1]
			}
		}
		for i := 1; i < rows; i++ {
			for j := 0; j < cols; j++ {
				matrix[i][j] += matrix[i-1][j]
			}
		}
	}
	return NumMatrix{matrix}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 < 0 {
		row1 = 0
	}
	if col1 < 0 {
		col1 = 0
	}
	if row2 > len(this.Matrix)-1 {
		row2 = len(this.Matrix) - 1
	}
	if col2 > len(this.Matrix[0])-1 {
		col2 = len(this.Matrix[0]) - 1
	}

	res := this.Matrix[row2][col2]
	if row1 > 0 {
		res -= this.Matrix[row1-1][col2]
	}
	if col1 > 0 {
		res -= this.Matrix[row2][col1-1]
	}
	if row1 > 0 && col1 > 0 {
		res += this.Matrix[row1-1][col1-1]
	}
	return res
}

/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */

func main() {
	matrix := [][]int{{3, 0, 1, 4, 2},
		{5, 6, 3, 2, 1},
		{1, 2, 0, 1, 5},
		{4, 1, 0, 1, 7},
		{1, 0, 3, 0, 5}}
	nums := Constructor(matrix)
	fmt.Println(nums.SumRegion(2, 1, 4, 3))
	fmt.Println(nums.SumRegion(1, 1, 2, 2))
	fmt.Println(nums.SumRegion(1, 2, 2, 4))
}
