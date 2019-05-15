package main

import (
	"fmt"
	"strings"
)

func recSol(s string, idx int) (string, int) {
	if !(s[idx] >= '0' && s[idx] <= '9') {
		return string(s[idx]), idx + 1
	}
	times := 0
	nowAt := idx
	for s[nowAt] != '[' {
		times = times*10 + int(s[nowAt]-'0')
		nowAt++
	}
	level := 1
	oneStr := ""
	nowAt++
	for nowAt < len(s) {
		if level == 0 {
			break
		}
		if s[nowAt] == ']' {
			level--
			nowAt++
		} else if s[nowAt] >= '0' && s[nowAt] <= '9' {
			tempStr, tempIdx := recSol(s, nowAt)
			oneStr += tempStr
			nowAt = tempIdx
		} else {
			oneStr += string(s[nowAt])
			nowAt++
		}
	}
	return strings.Repeat(oneStr, times), nowAt
}

func decodeString(s string) string {
	res := ""
	for i := 0; i < len(s); {
		temp, nextI := recSol(s, i)
		i = nextI
		res += temp
	}
	return res
}

func main() {
	s := "3[a]2[bc]"
	fmt.Println(decodeString(s))

	s = "3[a2[c]]"
	fmt.Println(decodeString(s))

	s = "2[abc]3[cd]ef"
	fmt.Println(decodeString(s))

	s = "3[a2[c]d]"
	fmt.Println(decodeString(s))

	s = "10[a]"
	fmt.Println(decodeString(s))
}
