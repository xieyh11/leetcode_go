package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	cache := make(map[string]bool, len(s))
	res := make([]string, 0)
	for i := 0; i <= len(s)-10; i++ {
		tempStr := s[i : 10+i]
		if v, ok := cache[tempStr]; ok {
			if v {
				res = append(res, tempStr)
				cache[tempStr] = false
			}
		} else {
			cache[tempStr] = true
		}
	}
	return res
}

func main() {
	s := "AAAAAAAAAAAA"
	fmt.Println(findRepeatedDnaSequences(s))
}
