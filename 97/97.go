package main

import (
	"fmt"
)

func dpSol(s1, s2, s3 string, i1, i2, i3 int) bool {
	if i1 >= len(s1) {
		return s3[i3:] == s2[i2:]
	}
	if i2 >= len(s2) {
		return s3[i3:] == s1[i1:]
	}
	if i3 >= len(s3) {
		return false
	}
	if s1[i1] == s3[i3] {
		if dpSol(s1, s2, s3, i1+1, i2, i3+1) {
			return true
		}
	}
	if s2[i2] == s3[i3] {
		if dpSol(s1, s2, s3, i1, i2+1, i3+1) {
			return true
		}
	}
	return false
}

func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s3) != len(s1)+len(s2) {
		return false
	}
	return dpSol(s1, s2, s3, 0, 0, 0)
}

func main() {
	s1 := "aabcc"
	s2 := "dbbca"
	s3 := "aadbbbaccc"
	fmt.Println(isInterleave(s1, s2, s3))
}
