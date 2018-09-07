package main

import (
	"fmt"
)

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}
	rows, cols := len(matrix), len(matrix[0])
	noZeroR, noZeroC := -1, -1
	for r := 0; r < rows; r++ {
		noZero := true
		for c := 0; c < cols; c++ {
			if matrix[r][c] == 0 {
				noZero = false
				break
			}
		}
		if noZero {
			noZeroR = r
			break
		}
	}
	for c := 0; c < cols; c++ {
		noZero := true
		for r := 0; r < rows; r++ {
			if matrix[r][c] == 0 {
				noZero = false
				break
			}
		}
		if noZero {
			noZeroC = c
			break
		}
	}
	if noZeroR >= 0 && noZeroC >= 0 {
		for r := 0; r < rows; r++ {
			if r != noZeroR {
				for c := 0; c < cols; c++ {
					if c != noZeroC {
						if matrix[r][c] == 0 {
							matrix[r][noZeroC] = 0
							matrix[noZeroR][c] = 0
						}
					}
				}
			}
		}
		for c := 0; c < cols; c++ {
			if c != noZeroC && matrix[noZeroR][c] == 0 {
				for r := 0; r < rows; r++ {
					if r != noZeroR {
						matrix[r][c] = 0
					}
				}
			}
		}
		for r := 0; r < rows; r++ {
			if r != noZeroR && matrix[r][noZeroC] == 0 {
				for c := 0; c < cols; c++ {
					if c != noZeroC {
						matrix[r][c] = 0
					}
				}
			}
		}
	} else {
		for r := 0; r < rows; r++ {
			for c := 0; c < cols; c++ {
				matrix[r][c] = 0
			}
		}
	}
}

func main() {
	matrix := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
	setZeroes(matrix)
	fmt.Println(matrix)
}
