package main

import (
	"fmt"
)

func isLive(val int) bool {
	return val == 1 || val == 2
}

func countNeighborLive(board [][]int, row, col int) int {
	liveCount := 0
	hasTop := row > 0
	hasBottom := row < len(board)-1
	hasLeft := col > 0
	hasRight := col < len(board[0])-1
	if hasTop {
		if hasLeft && isLive(board[row-1][col-1]) {
			liveCount++
		}
		if isLive(board[row-1][col]) {
			liveCount++
		}
		if hasRight && isLive(board[row-1][col+1]) {
			liveCount++
		}
	}
	if hasLeft && isLive(board[row][col-1]) {
		liveCount++
	}
	if hasRight && isLive(board[row][col+1]) {
		liveCount++
	}
	if hasBottom {
		if hasLeft && isLive(board[row+1][col-1]) {
			liveCount++
		}
		if isLive(board[row+1][col]) {
			liveCount++
		}
		if hasRight && isLive(board[row+1][col+1]) {
			liveCount++
		}
	}
	return liveCount
}

func gameOfLife(board [][]int) {
	//2 live to die, 3 die to live
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	for i := range board {
		for j := range board[i] {
			liveCount := countNeighborLive(board, i, j)
			if board[i][j] == 1 {
				if liveCount < 2 || liveCount > 3 {
					board[i][j] = 2
				}
			} else {
				if liveCount == 3 {
					board[i][j] = 3
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 2 {
				board[i][j] = 0
			} else if board[i][j] == 3 {
				board[i][j] = 1
			}
		}
	}
}

func main() {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	fmt.Println(board)
}
