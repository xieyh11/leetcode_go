package main

import (
	"container/list"
	"fmt"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	edges := make([]map[int]bool, numCourses)
	inDegrees := make([]int, numCourses)
	zeroInIdx := list.New()
	for i := range edges {
		edges[i] = make(map[int]bool)
	}
	for _, pair := range prerequisites {
		from, to := pair[1], pair[0]
		if _, ok := edges[from][to]; !ok {
			edges[from][to] = true
			inDegrees[to]++
		}
	}
	for i := range inDegrees {
		if inDegrees[i] == 0 {
			zeroInIdx.PushBack(i)
		}
	}
	res := []int{}
	for zeroInIdx.Len() > 0 {
		frontIdx := zeroInIdx.Remove(zeroInIdx.Front()).(int)
		res = append(res, frontIdx)
		for k := range edges[frontIdx] {
			inDegrees[k]--
			if inDegrees[k] == 0 {
				zeroInIdx.PushBack(k)
			}
		}
	}
	if len(res) < numCourses {
		return []int{}
	}
	return res
}

func main() {
	numCourses := 4
	pre := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	fmt.Println(findOrder(numCourses, pre))
}
