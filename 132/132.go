package main

import (
	"fmt"
)

func isPalin(s string) bool {
	if len(s) == 0 {
		return true
	}
	if len(s) == 1 {
		return true
	}
	for i, j := 0, len(s)-1; j >= i; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func minCut(s string) int {
	if isPalin(s) {
		return 0
	}
	res := make([]int, len(s))
	res[len(s)-1] = 0
	for i := len(s) - 2; i >= 0; i-- {
		if isPalin(s[i:]) {
			res[i] = 0
		} else {
			res[i] = res[i+1] + 1
			for j := i + 1; j < len(s)-1; j++ {
				if isPalin(s[i : j+1]) {
					if res[j+1]+1 < res[i] {
						res[i] = res[j+1] + 1
					}
				}
			}
		}
	}
	return res[0]
}

func main() {
	s := "aab"
	fmt.Println(minCut(s))
}
