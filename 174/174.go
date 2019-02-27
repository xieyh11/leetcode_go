package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 1
	}
	rows := len(dungeon)
	cols := len(dungeon[0])
	dungeon[rows-1][cols-1] = getMax(1, 1-dungeon[rows-1][cols-1])
	for c := cols - 2; c >= 0; c-- {
		dungeon[rows-1][c] = getMax(dungeon[rows-1][c+1]-dungeon[rows-1][c], 1)
	}
	for r := rows - 2; r >= 0; r-- {
		dungeon[r][cols-1] = getMax(dungeon[r+1][cols-1]-dungeon[r][cols-1], 1)
	}
	for r, c := rows-2, cols-2; r >= 0 && c >= 0; r, c = r-1, c-1 {
		for i := c; i >= 0; i-- {
			dungeon[r][i] = getMax(getMin(dungeon[r+1][i], dungeon[r][i+1])-dungeon[r][i], 1)
		}
		for j := r - 1; j >= 0; j-- {
			dungeon[j][c] = getMax(getMin(dungeon[j+1][c], dungeon[j][c+1])-dungeon[j][c], 1)
		}
	}
	return dungeon[0][0]
}

func main() {
	dungeon := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	fmt.Println(calculateMinimumHP(dungeon))
}
