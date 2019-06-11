package main

import (
	"fmt"
)

func isAnagram(a, b [26]int) bool {
	for i := 0; i < 26; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	pLetter := [26]int{}
	for i := range p {
		pLetter[p[i]-'a']++
	}
	sLetter := [26]int{}
	pLen := len(p)
	sLen := len(s)
	for i := 0; i < pLen; i++ {
		sLetter[s[i]-'a']++
	}
	res := []int{}
	for is, i := 0, pLen-1; i < sLen; is, i = is+1, i+1 {
		if isAnagram(sLetter, pLetter) {
			res = append(res, is)
		}
		if i+1 < sLen {
			sLetter[s[i+1]-'a']++
			sLetter[s[is]-'a']--
		}
	}
	return res
}

func main() {
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams(s, p))

	s = "abab"
	p = "ab"
	fmt.Println(findAnagrams(s, p))

	s = "abab"
	p = ""
	fmt.Println(findAnagrams(s, p))

	s = ""
	p = ""
	fmt.Println(findAnagrams(s, p))
}
