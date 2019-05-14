package main

import (
	"fmt"
)

func firstUniqChar(s string) int {
	letter := [26]int{}
	res := -1
	for i := range s {
		letter[s[i]-'a']++
	}
	for i := range s {
		if letter[s[i]-'a'] == 1 {
			res = i
			break
		}
	}
	return res
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))

	s = "loveleetcode"
	fmt.Println(firstUniqChar(s))
}
