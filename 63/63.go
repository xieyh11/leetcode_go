package main

import (
	"fmt"
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	rows := len(obstacleGrid)
	cols := len(obstacleGrid[0])
	if obstacleGrid[rows-1][cols-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}

	for i := 0; i < rows-1; i++ {
		for j := 0; j < cols-1; j++ {
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else {
				obstacleGrid[i][j] = -1
			}
		}
	}
	obstacleGrid[rows-1][cols-1] = 1
	for j := cols - 2; j >= 0; j-- {
		if obstacleGrid[rows-1][j] != 1 {
			obstacleGrid[rows-1][j] = obstacleGrid[rows-1][j+1]
		} else {
			obstacleGrid[rows-1][j] = 0
		}
	}
	for i := rows - 2; i >= 0; i-- {
		if obstacleGrid[i][cols-1] != 1 {
			obstacleGrid[i][cols-1] = obstacleGrid[i+1][cols-1]
		} else {
			obstacleGrid[i][cols-1] = 0
		}
	}

	for r, c := rows-2, cols-2; r >= 0 && c >= 0; r, c = r-1, c-1 {
		if obstacleGrid[r][c] == -1 {
			obstacleGrid[r][c] = obstacleGrid[r+1][c] + obstacleGrid[r][c+1]
		}
		for i := r - 1; i >= 0; i-- {
			if obstacleGrid[i][c] == -1 {
				obstacleGrid[i][c] = obstacleGrid[i+1][c] + obstacleGrid[i][c+1]
			}
		}
		for j := c - 1; j >= 0; j-- {
			if obstacleGrid[r][j] == -1 {
				obstacleGrid[r][j] = obstacleGrid[r+1][j] + obstacleGrid[r][j+1]
			}
		}
	}

	return obstacleGrid[0][0]
}

func main() {
	obstacle := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(obstacle))
}
