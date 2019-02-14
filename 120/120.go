package main

import (
	"fmt"
)

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	currentRes := make([]int, len(triangle))
	formerRes := make([]int, len(triangle)-1)
	currentRes[0] = triangle[1][0] + triangle[0][0]
	currentRes[1] = triangle[1][1] + triangle[0][0]
	for i := 2; i < len(triangle); i++ {
		copy(formerRes, currentRes[:i])
		currentRes[0] = triangle[i][0] + formerRes[0]
		currentRes[i] = triangle[i][i] + formerRes[i-1]
		for j := 1; j < i; j++ {
			minV := formerRes[j-1]
			if minV > formerRes[j] {
				minV = formerRes[j]
			}
			currentRes[j] = minV + triangle[i][j]
		}
	}
	minV := currentRes[0]
	for i := 1; i < len(triangle); i++ {
		if minV > currentRes[i] {
			minV = currentRes[i]
		}
	}
	return minV
}

func main() {
	triangle := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	fmt.Println(minimumTotal(triangle))
}
