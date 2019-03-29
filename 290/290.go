package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strs := strings.Split(str, " ")

	if len(pattern) != len(strs) {
		return false
	}

	match := [26]string{}
	reverseMatch := make(map[string]byte)

	for i := range pattern {
		idx := pattern[i] - 'a'
		if v, ok := reverseMatch[strs[i]]; !ok {
			if match[idx] != "" {
				return false
			}
			match[idx] = strs[i]
			reverseMatch[strs[i]] = pattern[i]
		} else if match[idx] == "" || pattern[i] != v {
			return false
		}
	}
	return true
}

func main() {
	pattren := "abba"
	strs := []string{"dog cat cat dog", "dog cat cat fish", "dog dog dog dog"}
	for _, str := range strs {
		fmt.Println(wordPattern(pattren, str))
	}
	fmt.Println(wordPattern("aaaa", strs[0]))
}
