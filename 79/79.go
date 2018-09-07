package main

import (
	"fmt"
)

func recurSearch(board [][]byte, r, c int, used [][]bool, word string) bool {
	if len(word) == 0 {
		return true
	}
	//up
	res := false
	if r > 0 && !used[r-1][c] && board[r-1][c] == word[0] {
		used[r-1][c] = true
		res = recurSearch(board, r-1, c, used, word[1:])
		used[r-1][c] = false
	}

	//down
	if !res && r < len(board)-1 && !used[r+1][c] && board[r+1][c] == word[0] {
		used[r+1][c] = true
		res = recurSearch(board, r+1, c, used, word[1:])
		used[r+1][c] = false
	}

	//left
	if !res && c > 0 && !used[r][c-1] && board[r][c-1] == word[0] {
		used[r][c-1] = true
		res = recurSearch(board, r, c-1, used, word[1:])
		used[r][c-1] = false
	}

	//right
	if !res && c < len(board[0])-1 && !used[r][c+1] && board[r][c+1] == word[0] {
		used[r][c+1] = true
		res = recurSearch(board, r, c+1, used, word[1:])
		used[r][c+1] = false
	}
	return res
}

func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	rows, cols := len(board), len(board[0])
	used := make([][]bool, rows)
	for i := range used {
		used[i] = make([]bool, cols)
	}
	res := false
	for r, row := range board {
		for c, ele := range row {
			if ele == word[0] {
				used[r][c] = true
				res = recurSearch(board, r, c, used, word[1:])
				used[r][c] = false
			}
			if res {
				break
			}
		}
		if res {
			break
		}
	}
	return res
}

func main() {
	board := [][]byte{{'A', 'A'}}
	word := "AAA"
	fmt.Println(exist(board, word))
}
