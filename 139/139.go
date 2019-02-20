package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	if len(s) == 0 {
		return true
	}
	if len(wordDict) == 0 {
		return false
	}
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	res := make([]bool, len(s))
	_, res[len(s)-1] = wordMap[s[len(s)-1:]]
	for i := len(s) - 2; i >= 0; i-- {
		if _, ok := wordMap[s[i:]]; ok {
			res[i] = true
		} else {
			tempRes := false
			for j := i + 1; j < len(s) && !tempRes; j++ {
				if _, ok1 := wordMap[s[i:j]]; ok1 {
					tempRes = res[j]
				}
			}
			res[i] = tempRes
		}
	}
	return res[0]
}

func main() {
	s := "abcd"
	wordDict := []string{"a", "abc", "b", "cd"}
	fmt.Println(wordBreak(s, wordDict))
}
