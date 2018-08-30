package main

import (
	"fmt"
)

func longestValidParentheses(s string) int {
	countEach := make([]int, len(s)+1)
	for i, v := range s {
		if v == '(' {
			countEach[i+1] = countEach[i] + 1
		} else {
			if countEach[i] > 0 {
				countEach[i+1] = countEach[i] - 1
			} else {
				countEach[i+1] = countEach[i]
			}
		}
	}
	if countEach[len(s)] > 0 {
		shouldMinus := countEach[len(s)]
		for i := len(s); i >= 0 && countEach[i] > 0; i-- {
			countEach[i] -= shouldMinus
		}
	}
	maxLen := 0
	currentLen := 0
	for i := 0; i < len(s); {
		if countEach[i+1] <= 0 {
			i++
			if currentLen != 0 {
				if currentLen > maxLen {
					maxLen = currentLen
				}
				currentLen = 0
			}
		} else {
			j := i + 1
			for ; j <= len(s) && countEach[j] > 0; j++ {
			}
			currentLen += j - i
			i = j
		}
	}
	if currentLen != 0 {
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	//From right
	countEach = make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			countEach[i] = countEach[i+1] + 1
		} else {
			if countEach[i+1] > 0 {
				countEach[i] = countEach[i+1] - 1
			} else {
				countEach[i] = countEach[i+1]
			}
		}
	}
	if countEach[0] > 0 {
		shouldMinus := countEach[0]
		for i := 0; i <= len(s) && countEach[i] > 0; i++ {
			countEach[i] -= shouldMinus
		}
	}
	currentLen = 0
	for i := len(s); i > 0; {
		if countEach[i-1] <= 0 {
			i--
			if currentLen != 0 {
				if currentLen > maxLen {
					maxLen = currentLen
				}
				currentLen = 0
			}
		} else {
			j := i - 1
			for ; j >= 0 && countEach[j] > 0; j-- {
			}
			currentLen += i - j
			i = j
		}
	}
	if currentLen != 0 {
		if currentLen > maxLen {
			maxLen = currentLen
		}
	}
	return maxLen
}

func main() {
	s := ")(((((()())()()))()(()))("
	fmt.Println(longestValidParentheses(s))
}
