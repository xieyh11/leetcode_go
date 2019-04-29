package main

import (
	"fmt"
)

func lessEnv(a, b []int) bool {
	return a[0] < b[0] && a[1] < b[1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}
	if len(envelopes) == 1 {
		return 1
	}
	inEdgeCount := make([]int, len(envelopes))
	toEdges := make([][]int, len(envelopes))
	for i := range envelopes {
		for j := i + 1; j < len(envelopes); j++ {
			if lessEnv(envelopes[i], envelopes[j]) {
				toEdges[j] = append(toEdges[j], i)
				inEdgeCount[i]++
			} else if lessEnv(envelopes[j], envelopes[i]) {
				toEdges[i] = append(toEdges[i], j)
				inEdgeCount[j]++
			}
		}
	}
	zeroIn := make([]int, 0)
	for i := range inEdgeCount {
		if inEdgeCount[i] == 0 {
			zeroIn = append(zeroIn, i)
		}
	}
	res := make([]int, len(envelopes))
	for len(zeroIn) > 0 {
		firstIdx := zeroIn[0]
		if res[firstIdx] == 0 {
			res[firstIdx] = 1
		}
		for _, toIdx := range toEdges[firstIdx] {
			res[toIdx] = getMax(res[toIdx], res[firstIdx]+1)
			inEdgeCount[toIdx]--
			if inEdgeCount[toIdx] == 0 {
				zeroIn = append(zeroIn, toIdx)
			}
		}
		zeroIn = zeroIn[1:]
	}
	maxRes := 0
	for _, num := range res {
		maxRes = getMax(maxRes, num)
	}
	return maxRes
}

func main() {
	env0 := [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}
	fmt.Println(maxEnvelopes(env0))
}
