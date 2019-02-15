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

func isEdge(edge []int, target int) bool {
	for _, k := range edge {
		if target == k {
			return true
		}
	}
	return false
}

func getByBfs(levels [][]int, edges [][]int, toIdx int) [][]int {
	if len(levels) == 0 {
		return [][]int{}
	}
	if len(levels) == 1 {
		res := make([][]int, 0)
		for _, nowIdx := range levels[0] {
			if isEdge(edges[nowIdx], toIdx) {
				res = append(res, []int{nowIdx})
			}
		}
		return res
	}

	res := make([][]int, 0)
	levelMinusOne := levels[:len(levels)-1]
	for _, nowIdx := range levels[len(levels)-1] {
		if isEdge(edges[nowIdx], toIdx) {
			tempRes := getByBfs(levelMinusOne, edges, nowIdx)
			for _, tempOne := range tempRes {
				tempOne = append(tempOne, nowIdx)
				res = append(res, tempOne)
			}
		}
	}
	return res
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	if len(wordList) == 0 {
		return [][]string{}
	}
	if len(beginWord) != len(endWord) || len(beginWord) != len(wordList[0]) {
		return [][]string{}
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
		return [][]string{}
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
		return [][]string{}
	}

	if len(edge[endIndex]) == 0 {
		return [][]string{}
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
		return [][]string{}
	}

	resInt := getByBfs(bfsLevel[:len(bfsLevel)-1], edge, endIndex)
	res := make([][]string, len(resInt))
	for i := range resInt {
		res[i] = make([]string, len(resInt[i])+1)
		res[i][0] = beginWord
		for j := 1; j < len(resInt[i]); j++ {
			res[i][j] = wordList[resInt[i][j]]
		}
		res[i][len(res[i])-1] = endWord
	}
	return res
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(findLadders(beginWord, endWord, wordList))
}
