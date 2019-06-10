package main

import (
	"fmt"
)

func countDiff(s1, s2 string) int {
	if len(s1) != len(s2) {
		return -1
	}
	count := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			count++
		}
	}
	return count
}

func minMutation(start string, end string, bank []string) int {
	if len(start) != len(end) {
		return -1
	}
	strIdx := make(map[string]int)
	count := 0
	strIdx[start] = count
	count++
	edges := make([][]int, count)
	for _, str := range bank {
		if _, ok := strIdx[str]; !ok {
			tempEdge := make([]int, 0)
			for k, idx := range strIdx {
				if countDiff(k, str) == 1 {
					tempEdge = append(tempEdge, idx)
					edges[idx] = append(edges[idx], count)
				}
			}
			strIdx[str] = count
			edges = append(edges, tempEdge)
			count++
		}
	}
	target, targetOk := strIdx[end]
	if !targetOk {
		return -1
	}

	if target == 0 {
		return 0
	}
	level := 1
	levelScan := []int{0}
	searched := make([]bool, count)
	searched[0] = true
	for len(levelScan) > 0 {
		nextScan := []int{}
		for _, idx := range levelScan {
			for _, to := range edges[idx] {
				if to == target {
					return level
				}
				if !searched[to] {
					searched[to] = true
					nextScan = append(nextScan, to)
				}
			}
		}
		levelScan = nextScan
		level++
	}

	return -1
}

func main() {
	s := "AACCGGTT"
	e := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(minMutation(s, e, bank))

	s = "AACCGGTT"
	e = "AAACGGTA"
	bank = []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}
	fmt.Println(minMutation(s, e, bank))

	s = "AAAAACCC"
	e = "AACCCCCC"
	bank = []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}
	fmt.Println(minMutation(s, e, bank))
}
