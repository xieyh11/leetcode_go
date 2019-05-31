package main

import (
	"fmt"
)

const PacificFlag = 0x1
const AtlanticFlag = 0x2
const BothFlag = PacificFlag | AtlanticFlag

func flowScan(matrix [][]int, scan [][]int, label [][]int, flag int, rows, cols int) {
	for len(scan) > 0 {
		nextScan := [][]int{}
		for _, loc := range scan {
			i, j := loc[0], loc[1]
			if i > 0 && label[i-1][j]&flag == 0 && matrix[i-1][j] >= matrix[i][j] {
				label[i-1][j] |= flag
				nextScan = append(nextScan, []int{i - 1, j})
			}
			if i < rows-1 && label[i+1][j]&flag == 0 && matrix[i+1][j] >= matrix[i][j] {
				label[i+1][j] |= flag
				nextScan = append(nextScan, []int{i + 1, j})
			}
			if j > 0 && label[i][j-1]&flag == 0 && matrix[i][j-1] >= matrix[i][j] {
				label[i][j-1] |= flag
				nextScan = append(nextScan, []int{i, j - 1})
			}
			if j < cols-1 && label[i][j+1]&flag == 0 && matrix[i][j+1] >= matrix[i][j] {
				label[i][j+1] |= flag
				nextScan = append(nextScan, []int{i, j + 1})
			}
		}
		scan = nextScan
	}
}

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	rows, cols := len(matrix), len(matrix[0])
	label := make([][]int, rows)
	for i := range label {
		label[i] = make([]int, cols)
	}
	scanP := [][]int{}
	scanA := [][]int{}
	for j := 0; j < cols; j++ {
		label[0][j] |= PacificFlag
		scanP = append(scanP, []int{0, j})

		label[rows-1][j] |= AtlanticFlag
		scanA = append(scanA, []int{rows - 1, j})
	}
	for i := 0; i < rows; i++ {
		label[i][0] |= PacificFlag
		if i > 0 {
			scanP = append(scanP, []int{i, 0})
		}

		label[i][cols-1] |= AtlanticFlag
		if i < rows-1 {
			scanA = append(scanA, []int{i, cols - 1})
		}
	}

	flowScan(matrix, scanP, label, PacificFlag, rows, cols)
	flowScan(matrix, scanA, label, AtlanticFlag, rows, cols)

	res := [][]int{}
	for i := range label {
		for j, v := range label[i] {
			if (v & BothFlag) == BothFlag {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}

func main() {
	matrix := [][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}
	fmt.Println(pacificAtlantic(matrix))
}
