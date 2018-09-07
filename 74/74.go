package main

import (
	"fmt"
)

func searchRow(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	if len(matrix) == 1 {
		return searchCol(matrix[0], target)
	}
	if len(matrix) == 2 {
		if target < matrix[0][0] {
			return false
		} else if target < matrix[1][0] {
			return searchCol(matrix[0], target)
		} else {
			return searchCol(matrix[1], target)
		}
	}
	midR := len(matrix) / 2
	if target < matrix[midR][0] {
		return searchRow(matrix[:midR], target)
	} else if target < matrix[midR+1][0] {
		return searchCol(matrix[midR], target)
	} else {
		return searchRow(matrix[midR+1:], target)
	}
}

func searchCol(row []int, target int) bool {
	if len(row) == 0 {
		return false
	}
	if len(row) == 1 {
		return row[0] == target
	}
	if len(row) == 2 {
		return (row[0] == target) || (row[1] == target)
	}
	midI := len(row) / 2
	if row[midI] == target {
		return true
	} else if row[midI] > target {
		return searchCol(row[:midI], target)
	} else {
		return searchCol(row[midI+1:], target)
	}
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return searchRow(matrix, target)
}

func main() {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 50}}
	target := 13
	fmt.Println(searchMatrix(matrix, target))
}
