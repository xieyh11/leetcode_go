package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	for i := range t {
		if t[i] == s[0] {
			return isSubsequence(s[1:], t[i+1:])
		}
	}
	return false
}

func main() {
	s := "abc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))

	s = "axc"
	t = "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
