package main

import (
	"fmt"
)

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])

	for j := cols - 2; j >= 0; j-- {
		grid[rows-1][j] += grid[rows-1][j+1]
	}
	for i := rows - 2; i >= 0; i-- {
		grid[i][cols-1] += grid[i+1][cols-1]
	}

	for r, c := rows-2, cols-2; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if grid[r+1][c] > grid[r][c+1] {
			grid[r][c] += grid[r][c+1]
		} else {
			grid[r][c] += grid[r+1][c]
		}

		for i := r - 1; i >= 0; i-- {
			if grid[i+1][c] > grid[i][c+1] {
				grid[i][c] += grid[i][c+1]
			} else {
				grid[i][c] += grid[i+1][c]
			}
		}
		for j := c - 1; j >= 0; j-- {
			if grid[r][j+1] > grid[r+1][j] {
				grid[r][j] += grid[r+1][j]
			} else {
				grid[r][j] += grid[r][j+1]
			}
		}
	}
	return grid[0][0]
}

func main() {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(minPathSum(grid))
}
