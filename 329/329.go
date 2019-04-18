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

func recSol(matrix [][]int, solution [][]int, i, j int) int {
	if solution[i][j] > 0 {
		return solution[i][j]
	}
	rows, cols := len(matrix), len(matrix[0])
	maxLen := 0
	if i > 0 && matrix[i-1][j] > matrix[i][j] {
		maxLen = getMax(maxLen, recSol(matrix, solution, i-1, j))
	}
	if i+1 < rows && matrix[i+1][j] > matrix[i][j] {
		maxLen = getMax(maxLen, recSol(matrix, solution, i+1, j))
	}
	if j > 0 && matrix[i][j-1] > matrix[i][j] {
		maxLen = getMax(maxLen, recSol(matrix, solution, i, j-1))
	}
	if j+1 < cols && matrix[i][j+1] > matrix[i][j] {
		maxLen = getMax(maxLen, recSol(matrix, solution, i, j+1))
	}
	solution[i][j] = maxLen + 1
	return solution[i][j]
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows, cols := len(matrix), len(matrix[0])
	solution := make([][]int, rows)
	for i := range solution {
		solution[i] = make([]int, cols)
	}
	res := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			res = getMax(res, recSol(matrix, solution, i, j))
		}
	}
	return res
}

func main() {
	matrix := [][]int{{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, {19, 18, 17, 16, 15, 14, 13, 12, 11, 10}, {20, 21, 22, 23, 24, 25, 26, 27, 28, 29}, {39, 38, 37, 36, 35, 34, 33, 32, 31, 30}, {40, 41, 42, 43, 44, 45, 46, 47, 48, 49}, {59, 58, 57, 56, 55, 54, 53, 52, 51, 50}, {60, 61, 62, 63, 64, 65, 66, 67, 68, 69}, {79, 78, 77, 76, 75, 74, 73, 72, 71, 70}, {80, 81, 82, 83, 84, 85, 86, 87, 88, 89}, {99, 98, 97, 96, 95, 94, 93, 92, 91, 90}, {100, 101, 102, 103, 104, 105, 106, 107, 108, 109}, {119, 118, 117, 116, 115, 114, 113, 112, 111, 110}, {120, 121, 122, 123, 124, 125, 126, 127, 128, 129}, {139, 138, 137, 136, 135, 134, 133, 132, 131, 130}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0}}
	fmt.Println(longestIncreasingPath(matrix))
}
