package main

import (
	"fmt"
)

type Loc [2]int

func borderDetect(detected [][]bool, from Loc) {
	current := []Loc{from}
	former := []Loc{}
	rows, cols := len(detected), len(detected[0])
	for len(current) > 0 {
		former = current
		current = make([]Loc, 0)
		for _, oneLoc := range former {
			if !detected[oneLoc[0]][oneLoc[1]] {
				detected[oneLoc[0]][oneLoc[1]] = true
				if oneLoc[1] < cols-1 {
					current = append(current, Loc{oneLoc[0], oneLoc[1] + 1})
				}
				if oneLoc[1] > 0 {
					current = append(current, Loc{oneLoc[0], oneLoc[1] - 1})
				}
				if oneLoc[0] < rows-1 {
					current = append(current, Loc{oneLoc[0] + 1, oneLoc[1]})
				}
				if oneLoc[0] > 0 {
					current = append(current, Loc{oneLoc[0] - 1, oneLoc[1]})
				}
			}
		}
	}
}

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	if len(board[0]) == 0 {
		return
	}
	rows, cols := len(board), len(board[0])
	detected := make([][]bool, rows)
	for i := range board {
		detected[i] = make([]bool, cols)
		for j := range board[i] {
			detected[i][j] = board[i][j] == 'X'
		}
	}
	for j := 0; j < cols; j++ {
		if board[0][j] == 'O' {
			borderDetect(detected, Loc{0, j})
		}
		if board[rows-1][j] == 'O' {
			borderDetect(detected, Loc{rows - 1, j})
		}
	}
	for i := 1; i < rows-1; i++ {
		if board[i][0] == 'O' {
			borderDetect(detected, Loc{i, 0})
		}
		if board[i][cols-1] == 'O' {
			borderDetect(detected, Loc{i, cols - 1})
		}
	}
	for i := range detected {
		for j := range detected[i] {
			if !detected[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	board := [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}
	solve(board)
	fmt.Println(board)
}
