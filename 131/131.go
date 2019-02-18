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

func recSol(s string, store [][][]string, hasValue []bool) {
	if len(s) == 0 {
		store[len(s)-1] = [][]string{}
		hasValue[len(s)-1] = true
		return
	}
	if len(s) == 1 {
		store[len(s)-1] = [][]string{{s}}
		hasValue[len(s)-1] = true
		return
	}
	res := make([][]string, 0)
	for i := 0; i < len(s); i++ {
		if isPalin(s[:i+1]) {
			if i < len(s)-1 {
				if !hasValue[len(s)-2-i] {
					recSol(s[i+1:], store, hasValue)
				}
				for _, onePart := range store[len(s)-2-i] {
					res = append(res, append([]string{s[:i+1]}, onePart...))
				}
			} else {
				res = append(res, []string{s})
			}
		}
	}
	hasValue[len(s)-1] = true
	store[len(s)-1] = res
}

func partition(s string) [][]string {
	store := make([][][]string, len(s))
	hasValue := make([]bool, len(s))
	recSol(s, store, hasValue)
	return store[len(s)-1]
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
