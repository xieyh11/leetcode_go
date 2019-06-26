package main

import (
	"fmt"
)

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	if n == 1 {
		return false
	}
	for i := 1; i <= n/2; i++ {
		temp := s[:i]
		allSame := true
		for j := i; j < n; j += i {
			if j+i > n || s[j:j+i] != temp {
				allSame = false
				break
			}
		}
		if allSame {
			return true
		}
	}
	return false
}

func main() {
	s := "abab"
	fmt.Println(repeatedSubstringPattern(s))

	s = "aba"
	fmt.Println(repeatedSubstringPattern(s))

	s = "abcabcabcabc"
	fmt.Println(repeatedSubstringPattern(s))

	s = "aabaaba"
	fmt.Println(repeatedSubstringPattern(s))
}
