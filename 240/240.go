package main

import (
	"fmt"
)

func halfSearch(row []int, target int) (int, bool) {
	if len(row) == 0 {
		return -1, false
	}
	if len(row) == 1 {
		if row[0] > target {
			return -1, false
		}
		return 0, row[0] == target
	}
	if len(row) == 2 {
		if row[0] > target {
			return -1, false
		}
		if row[0] == target {
			return 0, true
		}
		if row[0] < target && row[1] > target {
			return 0, false
		}
		return 1, row[1] == target
	}
	mid := len(row) / 2
	if row[mid] == target {
		return mid, true
	} else if row[mid] > target {
		return halfSearch(row[:mid], target)
	} else {
		idx, isin := halfSearch(row[mid+1:], target)
		return idx + mid + 1, isin
	}
}

func halfSearchCol(col [][]int, idx, target int) (int, bool) {
	if len(col) == 0 {
		return -1, false
	}
	if len(col) == 1 {
		if col[0][idx] > target {
			return -1, false
		}
		return 0, col[0][idx] == target
	}
	if len(col) == 2 {
		if col[0][idx] > target {
			return -1, false
		}
		if col[0][idx] == target {
			return 0, true
		}
		if col[0][idx] < target && col[1][idx] > target {
			return 0, false
		}
		return 1, col[1][idx] == target
	}
	mid := len(col) / 2
	if col[mid][idx] == target {
		return mid, true
	} else if col[mid][idx] > target {
		return halfSearchCol(col[:mid], idx, target)
	} else {
		idx, isin := halfSearchCol(col[mid+1:], idx, target)
		return idx + mid + 1, isin
	}
}

func searchRec(matrix [][]int, target, rowS, rowE, colS, colE int) bool {
	if rowE < rowS || colE < colS {
		return false
	}
	if rowE == rowS {
		_, temp := halfSearch(matrix[rowS][colS:colE+1], target)
		return temp
	}
	if rowE == rowS+1 {
		_, temp := halfSearch(matrix[rowS][colS:colE+1], target)
		if !temp {
			_, temp = halfSearch(matrix[rowE][colS:colE+1], target)
		}
		return temp
	}
	if colE == colS {
		_, temp := halfSearchCol(matrix[rowS:rowE+1], colS, target)
		return temp
	}
	if colE == colS+1 {
		_, temp := halfSearchCol(matrix[rowS:rowE+1], colS, target)
		if !temp {
			_, temp = halfSearchCol(matrix[rowS:rowE+1], colE, target)
		}
		return temp
	}
	midRow := (rowS + rowE) / 2
	if matrix[midRow][colS] > target {
		return searchRec(matrix, target, rowS, midRow-1, colS, colE)
	} else if matrix[midRow][colE] < target {
		return searchRec(matrix, target, midRow+1, rowE, colS, colE)
	}
	midIdx, midFound := halfSearch(matrix[midRow][colS:colE+1], target)
	if !midFound {
		_, colFound := halfSearchCol(matrix[rowS:rowE+1], midIdx+colS, target)
		if !colFound {
			return searchRec(matrix, target, midRow+1, rowE, colS, colS+midIdx-1) || searchRec(matrix, target, rowS, midRow-1, colS+midIdx+1, colE)
		}
		return true
	}
	return true
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	return searchRec(matrix, target, 0, len(matrix)-1, 0, len(matrix[0])-1)
}

func main() {
	matrix := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(searchMatrix(matrix, 9))
}
