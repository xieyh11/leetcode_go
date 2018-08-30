package main

import (
	"fmt"
)

func isRowValid(row []byte) bool {
	used := make(map[byte]bool)
	for _, b := range row {
		if b != '.' {
			if _, ok := used[b]; ok {
				return false
			} else {
				used[b] = true
			}
		}
	}
	return true
}

func isColumnValid(board [][]byte, col int) bool {
	used := make(map[byte]bool)
	for i := 0; i < 9; i++ {
		b := board[i][col]
		if b != '.' {
			if _, ok := used[b]; ok {
				return false
			} else {
				used[b] = true
			}
		}
	}
	return true
}

func isSubBoxValid(board [][]byte, row, col int) bool {
	used := make(map[byte]bool)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b := board[3*row+i][3*col+j]
			if b != '.' {
				if _, ok := used[b]; ok {
					return false
				} else {
					used[b] = true
				}
			}
		}
	}
	return true
}

func isValidSudoku(board [][]byte) bool {
	//row
	for i := 0; i < 9; i++ {
		if !isRowValid(board[i]) {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if !isColumnValid(board, j) {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !isSubBoxValid(board, i, j) {
				return false
			}
		}
	}
	return true
}

func main() {
	board := [][]byte{
		[]byte{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	fmt.Println(isValidSudoku(board))
}
