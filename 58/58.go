package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}
	count := 0
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			count++
		} else {
			break
		}
	}
	return count
}

func main() {
	s := ""
	fmt.Println(lengthOfLastWord(s))
}
