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

func recSol(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	first := points[0]
	nextLevel := make([][]int, 0)
	for i := 0; i < len(points); i++ {
		if points[i][0] > first[1] {
			nextLevel = append(nextLevel, points[i])
		}
	}
	return 1 + recSol(nextLevel)
}

func findMinArrowShots(points [][]int) int {
	sort.Sort(Ints(points))
	return recSol(points)
}

func main() {
	points := [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}
	fmt.Println(findMinArrowShots(points))
}
