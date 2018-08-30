package main

import (
	"fmt"
	"strings"
)

type WordsUsed struct {
	Used      map[string]int
	UsedCount int
}

func substringLoc(s, sub string) []int {
	subLen := len(sub)
	res := make([]int, 0)
	count := 0
	for len(s) >= subLen {
		if idx := strings.Index(s, sub); idx != -1 {
			res = append(res, count+idx)
			s = s[idx+1:]
			count += idx + 1
		} else {
			break
		}
	}
	return res
}

func reEvaluate(idx int, sToWords [][]string, wordsLen int, used WordsUsed) bool {
	if used.UsedCount == wordsLen {
		return true
	}
	if idx >= len(sToWords) || sToWords[idx] == nil || len(sToWords[idx]) == 0 {
		return false
	}
	res := false
	for _, whichWord := range sToWords[idx] {
		if used.Used[whichWord] == 0 {
			continue
		}
		used.Used[whichWord]--
		used.UsedCount++
		if temp := reEvaluate(idx+len(whichWord), sToWords, wordsLen, used); temp {
			res = true
		}
		used.Used[whichWord]++
		used.UsedCount--
		if res {
			break
		}
	}
	return res
}

func findSubstring(s string, words []string) []int {
	wordsLen := len(words)
	sToWords := make([][]string, len(s))

	used := WordsUsed{make(map[string]int), 0}
	for _, v := range words {
		if _, ok := used.Used[v]; !ok {
			tempRes := substringLoc(s, v)
			for _, idx := range tempRes {
				sToWords[idx] = append(sToWords[idx], v)
			}
			used.Used[v] = 1
		} else {
			used.Used[v]++
		}
	}
	res := make([]int, 0)
	for i := range s {
		if sToWords[i] != nil && len(sToWords[i]) > 0 {
			if temp := reEvaluate(i, sToWords, wordsLen, used); temp {
				res = append(res, i)
			}
		}
	}
	return res
}

func main() {
	s := strings.Repeat("a", 10000)
	words := []string{"a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a", "a"}

	fmt.Println(findSubstring(s, words))
}
