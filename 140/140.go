package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) []string {
	if len(s) == 0 || len(wordDict) == 0 {
		return []string{}
	}
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}
	resExist := make([]bool, len(s))
	_, resExist[len(s)-1] = wordMap[s[len(s)-1:]]
	for i := len(s) - 2; i >= 0; i-- {
		if _, ok := wordMap[s[i:]]; ok {
			resExist[i] = true
		} else {
			tempRes := false
			for j := i + 1; j < len(s) && !tempRes; j++ {
				if _, ok1 := wordMap[s[i:j]]; ok1 {
					tempRes = resExist[j]
				}
			}
			resExist[i] = tempRes
		}
	}
	if resExist[0] {
		res := make([][]string, len(s))
		if _, ok := wordMap[s[len(s)-1:]]; ok {
			res[len(s)-1] = []string{s[len(s)-1:]}
		} else {
			res[len(s)-1] = make([]string, 0)
		}
		for i := len(s) - 2; i >= 0; i-- {
			res[i] = make([]string, 0)
			if _, ok := wordMap[s[i:]]; ok {
				res[i] = append(res[i], s[i:])
			}
			for j := i + 1; j < len(s); j++ {
				tempStr := s[i:j]
				if _, ok1 := wordMap[tempStr]; ok1 {
					for _, v := range res[j] {
						res[i] = append(res[i], tempStr+" "+v)
					}
				}
			}
		}
		return res[0]
	} else {
		return []string{}
	}
}

func main() {
	s := "catsanddog"
	wordDict := []string{"cat", "cats", "and", "sand", "dog"}
	fmt.Println(wordBreak(s, wordDict))
}
