package main

import (
	"fmt"
)

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	if rowIndex == 1 {
		return []int{1, 1}
	}
	formerRow := make([]int, rowIndex)
	currentRow := make([]int, rowIndex+1)
	currentRow[0] = 1
	currentRow[1] = 1
	for i := 2; i <= rowIndex; i++ {
		copy(formerRow, currentRow[:i])
		currentRow[0] = 1
		currentRow[i] = 1
		for j := 1; j < i; j++ {
			currentRow[j] = formerRow[j-1] + formerRow[j]
		}
	}
	return currentRow
}

func main() {
	fmt.Println(getRow(4))
}
