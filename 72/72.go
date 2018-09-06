package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}
	rows := len(word1)
	cols := len(word2)
	res := make([][]int, rows)
	for i := range res {
		res[i] = make([]int, cols)
	}

	if word1[rows-1] != word2[cols-1] {
		res[rows-1][cols-1] = 1
	}
	for j := cols - 2; j >= 0; j-- {
		res[rows-1][j] = res[rows-1][j+1] + 1
		if word1[rows-1] == word2[j] && res[rows-1][j] > cols-1-j {
			res[rows-1][j] = cols - 1 - j
		}
	}
	for i := rows - 2; i >= 0; i-- {
		res[i][cols-1] = res[i+1][cols-1] + 1
		if word1[i] == word2[cols-1] && res[i][cols-1] > rows-1-i {
			res[i][cols-1] = rows - 1 - i
		}
	}
	for r, c := rows-2, cols-2; r >= 0 && c >= 0; r, c = r-1, c-1 {
		res[r][c] = min(res[r+1][c]+1, res[r][c+1]+1)
		if word1[r] == word2[c] {
			res[r][c] = min(res[r][c], res[r+1][c+1])
		} else {
			res[r][c] = min(res[r][c], res[r+1][c+1]+1)
		}
		for j := c - 1; j >= 0; j-- {
			res[r][j] = min(res[r+1][j]+1, res[r][j+1]+1)
			if word1[r] == word2[j] {
				res[r][j] = min(res[r][j], res[r+1][j+1])
			} else {
				res[r][j] = min(res[r][j], res[r+1][j+1]+1)
			}
		}
		for i := r - 1; i >= 0; i-- {
			res[i][c] = min(res[i+1][c]+1, res[i][c+1]+1)
			if word1[i] == word2[c] {
				res[i][c] = min(res[i][c], res[i+1][c+1])
			} else {
				res[i][c] = min(res[i][c], res[i+1][c+1]+1)
			}
		}
	}
	return res[0][0]
}

func main() {
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}
