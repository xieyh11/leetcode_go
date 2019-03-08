package main

import (
	"fmt"
)

func changeVal(searched []bool, idx int, val bool) {
	searched[idx] = val
}

func detectCircle(idx int, edges [][]int, searched []bool, connected []bool) bool {
	if len(edges[idx]) == 0 {
		return true
	}

	searched[idx] = true
	connected[idx] = true
	defer changeVal(searched, idx, false)

	for _, to := range edges[idx] {
		if searched[to] {
			return false
		} else {
			if !detectCircle(to, edges, searched, connected) {
				return false
			}
		}
	}
	return true
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	searched := make([]bool, numCourses)
	connected := make([]bool, numCourses)
	edges := make([][]int, numCourses)
	for i := range prerequisites {
		edges[prerequisites[i][1]] = append(edges[prerequisites[i][1]], prerequisites[i][0])
	}
	for i := range connected {
		if !connected[i] {
			if !detectCircle(i, edges, searched, connected) {
				return false
			}
		}
	}
	return true
}

func main() {
	nums := 2
	pairs := [][]int{{1, 0}}
	fmt.Println(canFinish(nums, pairs))
}
