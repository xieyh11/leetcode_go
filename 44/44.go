package main

import (
	"fmt"
)

func matchPart(s, p string) bool {
	for i := range s {
		if p[i] != '?' && p[i] != s[i] {
			return false
		}
	}
	return true
}

func dpMatch(s string, p string, already map[string](map[string]bool)) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	if len(p) == 1 {
		if p[0] == '*' {
			return true
		} else if p[0] == '?' {
			return len(s) == 1
		} else {
			return len(p) == len(s) && p[0] == s[0]
		}
	}
	if len(s) == 0 {
		for _, v := range p {
			if v != '*' {
				return false
			}
		}
		return true
	}
	if v1, ok1 := already[s]; ok1 {
		if v2, ok2 := v1[p]; ok2 {
			return v2
		}
	}
	i := 0
	for ; i < len(p) && p[i] != '*'; i++ {
	}

	res := false
	if i == 0 {
		for j := 0; j < len(s); j++ {
			if dpMatch(s[j:], p[1:], already) {
				res = true
				break
			}
		}
	} else {
		if len(s) < i {
			res = false
		} else {
			res = matchPart(s[:i], p[:i]) && dpMatch(s[i:], p[i:], already)
		}
	}

	if _, ok1 := already[s]; ok1 {
		already[s][p] = res
	} else {
		already[s] = make(map[string]bool)
		already[s][p] = res
	}

	return res
}

func isMatch(s string, p string) bool {
	already := make(map[string](map[string]bool))
	return dpMatch(s, p, already)
}

func main() {
	s := "ho"
	p := "ho**"
	fmt.Println(isMatch(s, p))
}
