package main

import (
	"fmt"
	"strings"
)

func generateParenthesisRecur(lAndR [2]int) []string {
	res := make([]string, 0)
	if lAndR[0] > 0 {
		tmepRes := generateParenthesisRecur([2]int{lAndR[0] - 1, lAndR[1] + 1})
		for i, str := range tmepRes {
			tmepRes[i] = "(" + str
		}
		res = append(res, tmepRes...)
	}
	if lAndR[1] > 0 {
		if lAndR[0] == 0 {
			res = append(res, strings.Repeat(")", lAndR[1]))
		} else {
			tmepRes := generateParenthesisRecur([2]int{lAndR[0], lAndR[1] - 1})
			for i, str := range tmepRes {
				tmepRes[i] = ")" + str
			}
			res = append(res, tmepRes...)
		}
	}
	return res
}

func generateParenthesis(n int) []string {
	return generateParenthesisRecur([2]int{n, 0})
}

func main() {
	fmt.Println(generateParenthesis(3))
}
