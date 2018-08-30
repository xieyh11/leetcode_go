package main

import (
	"fmt"
)

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	rows := make([]int, len(matrix))
	columns := make([]int, len(matrix[0]))
	for i := 0; i < len(rows); i++ {
		rows[i] = i
	}
	for j := 0; j < len(columns); j++ {
		columns[j] = j
	}
	res := make([]int, 0)
	step := 0
	for len(rows) > 0 && len(columns) > 0 {
		switch step {
		case 0:
			for j := 0; j < len(columns); j++ {
				res = append(res, matrix[rows[0]][columns[j]])
			}
			rows = rows[1:]
			step++
		case 1:
			for i := 0; i < len(rows); i++ {
				res = append(res, matrix[rows[i]][columns[len(columns)-1]])
			}
			columns = columns[:len(columns)-1]
			step++
		case 2:
			for j := len(columns) - 1; j >= 0; j-- {
				res = append(res, matrix[rows[len(rows)-1]][columns[j]])
			}
			rows = rows[:len(rows)-1]
			step++
		case 3:
			for i := len(rows) - 1; i >= 0; i-- {
				res = append(res, matrix[rows[i]][columns[0]])
			}
			columns = columns[1:]
			step = 0
		default:
			break
		}
	}
	return res
}

func main() {
	matrix := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix))
}
