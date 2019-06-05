package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func characterReplacement(s string, k int) int {
	res := 0
	for i := 0; i < len(s); {
		j := i + 1
		tempK := k
		firstHit := 0
		firstFound := false
		for ; j < len(s) && (s[i] == s[j] || tempK > 0); j++ {
			if s[i] != s[j] {
				tempK--
				if !firstFound {
					firstHit = j
					firstFound = true
				}
			}
		}
		res = getMax(res, j-i)
		if firstFound {
			tempK = k
			if tempK >= firstHit-i {
				tempK -= firstHit - i
				j = firstHit + 1
				for ; j < len(s) && (s[firstHit] == s[j] || tempK > 0); j++ {
					if s[firstHit] != s[j] {
						tempK--
					}
				}
				res = getMax(res, j-i)
			} else {
				res = getMax(res, tempK+1)
			}
			i = firstHit
		} else {
			i = j
		}
	}
	return res
}

func main() {
	s := "ABAB"
	k := 2
	fmt.Println(characterReplacement(s, k))

	s = "AABABBA"
	k = 1
	fmt.Println(characterReplacement(s, k))

	s = "ABBB"
	k = 1
	fmt.Println(characterReplacement(s, k))
}
