package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {

	lenS := len(s)
	if lenS == 0 {
		return 0
	}
	res := make([]int, lenS)
	if lastInt, _ := strconv.ParseInt(s[lenS-1:], 10, 64); lastInt >= 1 && lastInt <= 26 {
		res[lenS-1] = 1
	}
	for i := lenS - 2; i >= 0; i-- {
		if s[i] != '0' {
			for j := i + 1; j < lenS; j++ {
				if lastInt, _ := strconv.ParseInt(s[i:j], 10, 64); lastInt >= 1 && lastInt <= 26 {
					res[i] += res[j]
				} else {
					break
				}
			}
			if lastInt, _ := strconv.ParseInt(s[i:], 10, 64); lastInt >= 1 && lastInt <= 26 {
				res[i]++
			}
		}
	}
	return res[0]
}

func main() {
	str := "206"
	fmt.Println(numDecodings(str))
}
