package main

import (
	"fmt"
)

func recurSol(board [][]byte, column []bool, idx int) [][]string {
	if len(board) == 0 {
		return [][]string{}
	}
	n := len(board)
	if idx == n {
		res := make([]string, n)
		for i := 0; i < n; i++ {
			res[i] = string(board[i])
		}
		return [][]string{res}
	}
	res := make([][]string, 0)
	for i := 0; i < n; i++ {
		if !column[i] {
			exists := false
			for r1, c1 := idx+1, i+1; r1 < n && c1 < n; r1, c1 = r1+1, c1+1 {
				if board[r1][c1] == 'Q' {
					exists = true
					break
				}
			}
			for r1, c1 := idx-1, i-1; !exists && r1 >= 0 && c1 >= 0; r1, c1 = r1-1, c1-1 {
				if board[r1][c1] == 'Q' {
					exists = true
					break
				}
			}
			for r1, c1 := idx-1, i+1; !exists && r1 >= 0 && c1 < n; r1, c1 = r1-1, c1+1 {
				if board[r1][c1] == 'Q' {
					exists = true
					break
				}
			}
			for r1, c1 := idx+1, i-1; !exists && r1 < n && c1 >= 0; r1, c1 = r1+1, c1-1 {
				if board[r1][c1] == 'Q' {
					exists = true
					break
				}
			}
			if !exists {
				board[idx][i] = 'Q'
				column[i] = true
				temp := recurSol(board, column, idx+1)
				res = append(res, temp...)
				board[idx][i] = '.'
				column[i] = false
			}
		}
	}
	return res
}

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	column := make([]bool, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	return recurSol(board, column, 0)
}

func main() {
	for i := 1; i < 12; i++ {
		fmt.Println(i, len(solveNQueens(i)))
	}
}
