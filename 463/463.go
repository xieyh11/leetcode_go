package main

import (
	"fmt"
)

func islandPerimeter(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	rows, cols := len(grid), len(grid[0])
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				res += 4
				if i > 0 && grid[i-1][j] == 1 {
					res--
				}
				if i < rows-1 && grid[i+1][j] == 1 {
					res--
				}
				if j > 0 && grid[i][j-1] == 1 {
					res--
				}
				if j < cols-1 && grid[i][j+1] == 1 {
					res--
				}
			}
		}
	}
	return res
}

func main() {
	grid := [][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}
	fmt.Println(islandPerimeter(grid))
}
