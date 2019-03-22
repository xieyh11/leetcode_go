package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	alphabets := [26]int{0}
	countNonZero := 0
	for i := range s {
		if alphabets[s[i]-'a'] == 0 {
			countNonZero++
		}
		alphabets[s[i]-'a']++
	}
	for i := range t {
		alphabets[t[i]-'a']--
		if alphabets[t[i]-'a'] == 0 {
			countNonZero--
		}
		if alphabets[t[i]-'a'] < 0 {
			return false
		}
	}
	return countNonZero == 0
}

func main() {
	s := "cat"
	t := "rat"
	fmt.Println(isAnagram(s, t))
}
