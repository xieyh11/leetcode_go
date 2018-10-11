package main

import (
	"fmt"
)

func byRows(matrix [][]int) int {
	rows, cols := len(matrix), len(matrix[0])
	matrixSum := make([][]int, rows)
	for i := range matrixSum {
		matrixSum[i] = make([]int, cols)
		for j := range matrixSum[i] {
			if i > 0 {
				matrixSum[i][j] = matrix[i][j] + matrixSum[i-1][j]
			} else {
				matrixSum[i][j] = matrix[i][j]
			}
		}
	}
	max := 0
	for startR := 0; startR < rows; startR++ {
		for endR := startR; endR < rows; endR++ {
			sumR := make([]int, cols)
			copy(sumR, matrixSum[endR])
			if startR > 0 {
				for i := range sumR {
					sumR[i] -= matrixSum[startR-1][i]
				}
			}
			maxL := 0
			target := endR - startR + 1
			for i := 0; i < cols; {
				if sumR[i] == target {
					j := i + 1
					for ; j < cols; j++ {
						if sumR[j] != target {
							break
						}
					}
					tempL := j - i
					if maxL < tempL {
						maxL = tempL
					}
					i = j
				} else {
					i++
				}
			}
			if max < maxL*target {
				max = maxL * target
			}
		}
	}
	return max
}

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	var matrixR [][]int
	if rows <= cols {
		matrixR = make([][]int, rows)
		for i := range matrixR {
			matrixR[i] = make([]int, cols)
			for j := range matrixR[i] {
				matrixR[i][j] = int(matrix[i][j] - '0')
			}
		}
	} else {
		matrixR = make([][]int, cols)
		for j := range matrixR {
			matrixR[j] = make([]int, rows)
			for i := range matrixR[j] {
				matrixR[j][i] = int(matrix[i][j] - '0')
			}
		}
	}
	return byRows(matrixR)
}

func main() {
	matrix := [][]byte{{'0', '0', '0', '1', '0', '1', '1', '1'}, {'0', '1', '1', '0', '0', '1', '0', '1'}, {'1', '0', '1', '1', '1', '1', '0', '1'}, {'0', '0', '0', '1', '0', '0', '0', '0'}, {'0', '0', '1', '0', '0', '0', '1', '0'}, {'1', '1', '1', '0', '0', '1', '1', '1'}, {'1', '0', '0', '1', '1', '0', '0', '1'}, {'0', '1', '0', '0', '1', '1', '0', '0'}, {'1', '0', '0', '1', '0', '0', '0', '0'}}
	fmt.Println(maximalRectangle(matrix))
}
