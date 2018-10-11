package main

import (
	"fmt"
)

func isScramble(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	if len(s1) == 0 {
		return true
	}
	lenS := len(s1)
	res := make([][][]bool, lenS)
	res[0] = make([][]bool, lenS)
	for i := range s1 {
		res[0][i] = make([]bool, lenS)
		for j := range s2 {
			res[0][i][j] = (s1[i] == s2[j])
		}
	}
	for n := 1; n < lenS; n++ {
		res[n] = make([][]bool, lenS)
		for i := 0; i < lenS-n; i++ {
			res[n][i] = make([]bool, lenS)
			for j := 0; j < lenS-n; j++ {
				temp := false
				for k := 0; k < n; k++ {
					rK := n - k - 1
					temp = temp || (res[k][i][j] && res[rK][i+k+1][j+k+1]) || (res[k][i][j+rK+1] && res[rK][i+k+1][j])
					if temp {
						break
					}
				}
				res[n][i][j] = temp
			}
		}
	}
	return res[lenS-1][0][0]
}

func main() {
	str1 := "abb"
	str2 := "bba"
	fmt.Println(isScramble(str1, str2))
}
