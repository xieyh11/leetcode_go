package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(words []string) int {
	wordsLen := len(words)
	if wordsLen == 0 {
		return 0
	}
	pairs := make([][]bool, wordsLen)
	for i := range pairs {
		pairs[i] = make([]bool, wordsLen)
	}
	letters := [26][]int{}
	for i := range words {
		for _, le := range words[i] {
			idx := le - 'a'
			if len(letters[idx]) == 0 || letters[idx][len(letters[idx])-1] != i {
				letters[idx] = append(letters[idx], i)
			}
		}
	}
	for _, oneLetter := range letters {
		for i := range oneLetter {
			for j := i + 1; j < len(oneLetter); j++ {
				pairs[oneLetter[i]][oneLetter[j]] = true
			}
		}
	}
	max := 0
	for i := range words {
		for j := i + 1; j < len(words); j++ {
			if !pairs[i][j] {
				max = getMax(max, len(words[i])*len(words[j]))
			}
		}
	}
	return max
}

func main() {
	a := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Println(maxProduct(a))
}
