package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s2t := make(map[byte]byte, len(s))
	t2s := make(map[byte]byte, len(s))
	for i := range s {
		vs, oks := s2t[s[i]]
		vt, okt := t2s[t[i]]
		if oks && okt {
			if vs != t[i] || vt != s[i] {
				return false
			}
		} else if !oks && !okt {
			t2s[t[i]] = s[i]
			s2t[s[i]] = t[i]
		} else {
			return false
		}
	}
	return true
}

func main() {
	s := "paper"
	t := "title"
	fmt.Println(isIsomorphic(s, t))
}
