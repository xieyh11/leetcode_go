package main

import (
	"fmt"
)

func countAll(sub string, subExist []int) {
	subLen := len(sub)
	for i := range sub {
		offset := 1 << uint(sub[i]-'a')
		for j := i; j < subLen; j++ {
			subExist[j-i] |= offset
		}
	}
}

func countOne(a int) int {
	count := 0
	for a != 0 {
		a &= (a - 1)
		count++
	}
	return count
}

func findSubstringInWraproundString(p string) int {
	pLen := len(p)
	if pLen == 0 {
		return 0
	}
	subExist := make([]int, pLen)
	left := 0
	for left < pLen {
		byteNow := (p[left]+1-'a')%26 + 'a'
		j := left + 1
		for j < pLen && p[j] == byteNow {
			j++
			byteNow = (byteNow+1-'a')%26 + 'a'
		}
		countAll(p[left:j], subExist)
		left = j
	}
	res := 0
	for i := range subExist {
		res += countOne(subExist[i])
	}
	return res
}

func main() {
	s := "a"
	fmt.Println(findSubstringInWraproundString(s))

	s = "zab"
	fmt.Println(findSubstringInWraproundString(s))
}
