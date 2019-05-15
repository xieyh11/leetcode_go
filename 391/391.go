package main

import (
	"fmt"
	"sort"
)

type SortInts [][]int

func (this SortInts) Len() int {
	return len(this)
}

func (this SortInts) Less(i, j int) bool {
	return this[i][1] < this[j][1]
}

func (this SortInts) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

type SortInts1 [][]int

func (this SortInts1) Len() int {
	return len(this)
}

func (this SortInts1) Less(i, j int) bool {
	return this[i][0] < this[j][0]
}

func (this SortInts1) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

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

func earseOut(scan [][]int, row int) [][]int {
	if len(scan) == 0 {
		return scan
	}
	left := 0
	right := len(scan) - 1
	for left < right {
		if scan[left][3] > row {
			left++
		} else if scan[right][3] <= row {
			right--
		} else {
			scan[left], scan[right] = scan[right], scan[left]
			left++
			right--
		}
	}
	if scan[left][3] > row {
		left++
	}
	return scan[:left]
}

func isRectangleCover(rectangles [][]int) bool {
	if len(rectangles) == 0 {
		return false
	}
	if len(rectangles) == 1 {
		return true
	}
	sort.Sort(SortInts(rectangles))
	rowStart, rowEnd, colStart, colEnd := rectangles[0][1], rectangles[0][3], rectangles[0][0], rectangles[0][2]
	for i := 1; i < len(rectangles); i++ {
		rowStart = getMin(rowStart, rectangles[i][1])
		rowEnd = getMax(rowEnd, rectangles[i][3])
		colStart = getMin(colStart, rectangles[i][0])
		colEnd = getMax(colEnd, rectangles[i][2])
	}

	scanList := [][]int{}
	scanStart := 0
	for row := rowStart; row < rowEnd; row++ {
		scanList = earseOut(scanList, row)
		for scanStart < len(rectangles) && rectangles[scanStart][1] <= row {
			scanList = append(scanList, rectangles[scanStart])
			scanStart++
		}
		col := colStart
		sort.Sort(SortInts1(scanList))
		for i := range scanList {
			if col == scanList[i][0] {
				col = scanList[i][2]
			} else {
				return false
			}
		}
		if col < colEnd {
			return false
		}
	}
	return true
}

func main() {
	matrix := [][]int{{1, 1, 3, 3},
		{3, 1, 4, 2},
		{3, 2, 4, 4},
		{1, 3, 2, 4},
		{2, 3, 3, 4}}
	fmt.Println(isRectangleCover(matrix))

	matrix = [][]int{{1, 1, 2, 3}, {1, 3, 2, 4}, {3, 1, 4, 2}, {3, 2, 4, 4}}
	fmt.Println(isRectangleCover(matrix))

	matrix = [][]int{{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{3, 2, 4, 4}}
	fmt.Println(isRectangleCover(matrix))

	matrix = [][]int{{1, 1, 3, 3},
		{3, 1, 4, 2},
		{1, 3, 2, 4},
		{2, 2, 4, 4}}
	fmt.Println(isRectangleCover(matrix))
}
