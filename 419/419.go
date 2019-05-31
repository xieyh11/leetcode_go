package main

import (
	"fmt"
)

func countBattleships(board [][]byte) int {
	count := 0
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'X' {
				if i > 0 && board[i-1][j] == 'X' {
					continue
				}
				if j > 0 && board[i][j-1] == 'X' {
					continue
				}
				count++
			}
		}
	}
	return count
}

func main() {
	board := [][]byte{[]byte("X..X"), []byte("...X"), []byte("...X")}
	fmt.Println(countBattleships(board))
}
