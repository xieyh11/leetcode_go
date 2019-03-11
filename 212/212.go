package main

import (
	"fmt"
)

func changeVal(searched [][]bool, loc [2]int, val bool) {
	searched[loc[0]][loc[1]] = val
}

func search(board [][]byte, word string, start [2]int, searched [][]bool) bool {
	if len(word) == 0 {
		return true
	}
	searched[start[0]][start[1]] = true
	defer changeVal(searched, start, false)

	rows, cols := len(board), len(board[0])
	find := false
	if start[1] > 0 && !searched[start[0]][start[1]-1] && board[start[0]][start[1]-1] == word[0] {
		find = search(board, word[1:], [2]int{start[0], start[1] - 1}, searched)
	}
	if !find && start[1] < cols-1 && !searched[start[0]][start[1]+1] && board[start[0]][start[1]+1] == word[0] {
		find = search(board, word[1:], [2]int{start[0], start[1] + 1}, searched)
	}
	if !find && start[0] > 0 && !searched[start[0]-1][start[1]] && board[start[0]-1][start[1]] == word[0] {
		find = search(board, word[1:], [2]int{start[0] - 1, start[1]}, searched)
	}
	if !find && start[0] < rows-1 && !searched[start[0]+1][start[1]] && board[start[0]+1][start[1]] == word[0] {
		find = search(board, word[1:], [2]int{start[0] + 1, start[1]}, searched)
	}
	return find
}

func findWords(board [][]byte, words []string) []string {
	if len(words) == 0 || len(board) == 0 || len(board[0]) == 0 {
		return []string{}
	}
	locs := [26][][2]int{}
	rows, cols := len(board), len(board[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			idx := board[i][j] - 'a'
			locs[idx] = append(locs[idx], [2]int{i, j})
		}
	}
	searched := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		searched[i] = make([]bool, cols)
	}
	res := make([]string, 0)
	searchWord := make(map[string]bool)
	for _, word := range words {
		if _, ok := searchWord[word]; !ok {
			if len(word) != 0 {
				idx := word[0] - 'a'
				for _, loc := range locs[idx] {
					if search(board, word[1:], loc, searched) {
						res = append(res, word)
						searchWord[word] = true
						break
					}
				}
			} else {
				res = append(res, word)
				searchWord[word] = true
			}
		}
	}
	return res
}

func main() {
	board := [][]byte{[]byte("oaan"), []byte("etae"), []byte("ihkr"), []byte("iflv")}
	words := []string{"oath", "pea", "eat", "rain"}
	fmt.Println(findWords(board, words))
}
