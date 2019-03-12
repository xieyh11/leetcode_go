package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// An implemention of Mancher Algorithm for longest palindrome
func ManacherAlgorithm(s string) []int {
	manacherLen := 2*len(s) + 1
	manacherByte := make([]byte, manacherLen)
	for i := 0; i < manacherLen; i += 2 {
		manacherByte[i] = '#'
	}
	for i := 1; i < manacherLen; i += 2 {
		manacherByte[i] = s[i/2]
	}
	res := make([]int, manacherLen)
	res[0] = 1
	mx := 1
	mid := 0
	for i := 1; i < manacherLen; i++ {
		if mx > i {
			res[i] = getMin(mx-i, res[2*mid-i])
		} else {
			res[i] = 1
		}
		for i-res[i] >= 0 && i+res[i] < manacherLen && manacherByte[i-res[i]] == manacherByte[i+res[i]] {
			res[i]++
		}
		if mx < res[i]+i {
			mx = res[i] + i
			mid = i
		}
	}
	return res
}

func reverseString(s string) string {
	sByte := []byte(s)
	for i, j := 0, len(sByte)-1; i < j; i, j = i+1, j-1 {
		sByte[i], sByte[j] = sByte[j], sByte[i]
	}
	return string(sByte)
}

func shortestPalindrome(s string) string {
	palin := ManacherAlgorithm(s)
	maxIdx := 0
	for i := len(palin) - 1; i >= 0; i-- {
		if palin[i]-1 == i {
			maxIdx = i
			break
		}
	}
	maxLen := palin[maxIdx] - 1
	suffix := s[maxLen:]
	return reverseString(suffix) + s
}

func main() {
	s := "a"
	fmt.Println(shortestPalindrome(s))
}
