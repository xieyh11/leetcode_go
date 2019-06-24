package main

import (
	"fmt"
	"sort"
)

type Ints [][]int

func (a Ints) Len() int {
	return len(a)
}

func (a Ints) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Ints) Less(i, j int) bool {
	return a[i][1] < a[j][1]
}

func recSol(points [][]int, former int) int {
	if len(points) == 0 {
		return 0
	}
	first := []int{}
	firstGet := -1
	for i := range points {
		if points[i][0] > former {
			first = points[i]
			firstGet = i
			break
		}
	}
	if firstGet == -1 {
		return 0
	}
	return 1 + recSol(points[firstGet+1:], first[1])
}

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	sort.Sort(Ints(points))
	return 1 + recSol(points[1:], points[0][1])
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
