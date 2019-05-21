package main

import (
	"fmt"
)

func longestPalindrome(s string) int {
	letters := make(map[byte]int)
	for i := range s {
		letters[s[i]]++
	}

	res := 0
	hasOdd := false
	lastZero := ^1
	for _, v := range letters {
		if v&1 == 1 {
			hasOdd = true
		}
		res += (v & lastZero)
	}
	if hasOdd {
		res++
	}
	return res
}

func main() {
	s := "abccccdd"
	fmt.Println(longestPalindrome(s))
}
