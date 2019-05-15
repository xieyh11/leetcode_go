package main

import (
	"fmt"
)

func findTheDifference(s string, t string) byte {
	letters := [26]int{}
	for i := range s {
		letters[s[i]-'a']++
	}
	res := byte(0)
	for i := range t {
		idx := t[i] - 'a'
		if letters[idx] == 0 {
			res = t[i]
			break
		}
		letters[idx]--
	}
	return res
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(string(findTheDifference(s, t)))
}
