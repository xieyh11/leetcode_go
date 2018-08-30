package main

import (
	"fmt"
)

func generateMatrix(n int) [][]int {
	rows := make([]int, n)
	columns := make([]int, n)
	for i := 0; i < n; i++ {
		rows[i] = i
	}
	copy(columns, rows)
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	step := 0
	count := 1
	for len(rows) > 0 && len(columns) > 0 {
		switch step {
		case 0:
			for j := 0; j < len(columns); j++ {
				matrix[rows[0]][columns[j]] = count
				count++
			}
			rows = rows[1:]
			step++
		case 1:
			for i := 0; i < len(rows); i++ {
				matrix[rows[i]][columns[len(columns)-1]] = count
				count++
			}
			columns = columns[:len(columns)-1]
			step++
		case 2:
			for j := len(columns) - 1; j >= 0; j-- {
				matrix[rows[len(rows)-1]][columns[j]] = count
				count++
			}
			rows = rows[:len(rows)-1]
			step++
		case 3:
			for i := len(rows) - 1; i >= 0; i-- {
				matrix[rows[i]][columns[0]] = count
				count++
			}
			columns = columns[1:]
			step = 0
		default:
			break
		}
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3))
}
