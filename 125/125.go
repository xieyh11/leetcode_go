package main

import (
	"fmt"
	"strings"
)

func isLetter(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z')
}

func isPalindrome(s string) bool {
	if s == "" {
		return true
	}

	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for right >= left {
		leftIs := isLetter(s[left])
		if !leftIs {
			left++
			continue
		}
		rightIs := isLetter(s[right])
		if !rightIs {
			right--
			continue
		}
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "race a car"
	fmt.Println(isPalindrome(s))
}
