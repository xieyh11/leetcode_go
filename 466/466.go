package main

import (
	"fmt"
)

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len*n1 < s2Len {
		return 0
	}
	count1 := 0
	count2 := 0
	s1Scan := 0
	s2Scan := 0
	subMatch := false
	inComplete := [][]int{}
	complete := []int{}
	for !subMatch && count1 < n1 {
		if s1[s1Scan] == s2[s2Scan] {
			s2Scan++
			if s2Scan == s2Len {
				s2Scan = 0
				count2++
			}
		}
		s1Scan++
		if s1Scan == s1Len {
			s1Scan = 0
			count1++
		}
		if s2Scan == 0 && count2 > 0 && count2%n2 == 0 {
			if s1Scan == 0 && count1 > 0 {
				complete = []int{count1, count2 / n2}
				subMatch = true
			} else {
				inComplete = append(inComplete, []int{count1 + 1, count2 / n2})
			}
		}
	}

	res := 0
	if subMatch {
		res = n1 / complete[0] * complete[1]
		remain := n1 % complete[0]
		for i := len(inComplete) - 1; i >= 0; i-- {
			if inComplete[i][0] <= remain {
				res += inComplete[i][1]
				break
			}
		}
	} else {
		for i := len(inComplete) - 1; i >= 0; i-- {
			if inComplete[i][0] <= n1 {
				res += inComplete[i][1]
				break
			}
		}
	}
	return res
}

func main() {
	s1, n1 := "acb", 4
	s2, n2 := "ab", 2
	fmt.Println(getMaxRepetitions(s1, n1, s2, n2))

	s1, n1 = "aaa", 3
	s2, n2 = "aa", 1
	fmt.Println(getMaxRepetitions(s1, n1, s2, n2))
}
