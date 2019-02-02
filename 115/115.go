package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	if len(s) == 0 {
		if len(t) == 0 {
			return 1
		} else {
			return 0
		}
	}
	res := make([][]int, len(t))
	for i := range res {
		res[i] = make([]int, len(s))
	}
	for r := len(t) - 1; r >= 0; r-- {
		for c := len(s) - 1; c >= 0; c-- {
			if t[r] == s[c] {
				if r == len(t)-1 && c == len(s)-1 {
					res[r][c] = 1
				} else if c == len(s)-1 {
					res[r][c] = 0
				} else if r == len(t)-1 {
					res[r][c] = res[r][c+1] + 1
				} else {
					res[r][c] = res[r][c+1] + res[r+1][c+1]
				}
			} else {
				if c == len(s)-1 {
					res[r][c] = 0
				} else {
					res[r][c] = res[r][c+1]
				}
			}
		}
	}
	return res[0][0]
}

func main() {
	s, t := "ddd", "dd"
	fmt.Println(numDistinct(s, t))
}
