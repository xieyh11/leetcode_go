package main

import (
	"fmt"
)

func isDifferByOne(s1, s2 string) bool {
	diffCount := 0
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			diffCount++
		}
	}
	return diffCount == 1
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(wordList) == 0 {
		return 0
	}
	if len(beginWord) != len(endWord) || len(beginWord) != len(wordList[0]) {
		return 0
	}
	endIsIn := false
	endIndex := len(wordList)
	for i, word := range wordList {
		if word == endWord {
			endIsIn = true
			endIndex = i
		}
	}
	if !endIsIn {
		return 0
	}

	wordListLen := len(wordList)

	edge := make([][]int, wordListLen+1) // len(wordList) is beginWord index
	for i := range edge {
		edge[i] = make([]int, 0)
	}
	for i := range wordList {
		for j := i + 1; j < wordListLen; j++ {
			if isDifferByOne(wordList[i], wordList[j]) {
				edge[i] = append(edge[i], j)
				edge[j] = append(edge[j], i)
			}
		}
	}
	for i := range wordList {
		if isDifferByOne(beginWord, wordList[i]) {
			edge[i] = append(edge[i], wordListLen)
			edge[wordListLen] = append(edge[wordListLen], i)
		}
	}

	if len(edge[wordListLen]) == 0 {
		return 0
	}

	if len(edge[endIndex]) == 0 {
		return 0
	}

	bfsLevel := make([][]int, 1)
	bfsLevel[0] = []int{wordListLen}

	searchCount := 0
	newLevel := true
	searched := make([]bool, wordListLen+1)
	searched[wordListLen] = true
	for searchCount < wordListLen && newLevel && !searched[endIndex] {
		tempLevel := make([]int, 0)
		for _, idx := range bfsLevel[len(bfsLevel)-1] {
			edgeOf := edge[wordListLen]
			if idx < wordListLen {
				edgeOf = edge[idx]
			}
			for _, idxTo := range edgeOf {
				if !searched[idxTo] {
					searched[idxTo] = true
					tempLevel = append(tempLevel, idxTo)
					searchCount++
				}
			}
		}
		if len(tempLevel) > 0 {
			bfsLevel = append(bfsLevel, tempLevel)
		} else {
			newLevel = false
		}
	}
	if !searched[endIndex] {
		return 0
	}
	return len(bfsLevel)
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength(beginWord, endWord, wordList))
}
