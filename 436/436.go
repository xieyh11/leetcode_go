package main

import (
	"fmt"
	"sort"
)

type LeftBound struct {
	Bound int
	Index int
}

type Bounds []LeftBound

func (a Bounds) Len() int {
	return len(a)
}

func (a Bounds) Less(i, j int) bool {
	return a[i].Bound < a[j].Bound
}

func (a Bounds) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func halfFind(bounds []LeftBound, target int) (int, bool) {
	if len(bounds) == 0 {
		return -1, false
	}
	if len(bounds) == 1 {
		if bounds[0].Bound > target {
			return -1, false
		}
		return 0, bounds[0].Bound == target
	}
	if len(bounds) == 2 {
		if bounds[0].Bound > target {
			return -1, false
		} else if bounds[1].Bound > target {
			return 0, bounds[0].Bound == target
		}
		return 1, bounds[1].Bound == target
	}

	mid := len(bounds) / 2
	if bounds[mid].Bound == target {
		return mid, true
	} else if bounds[mid].Bound > target {
		return halfFind(bounds[:mid], target)
	} else {
		idx, found := halfFind(bounds[mid+1:], target)
		return idx + mid + 1, found
	}
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	if n == 0 {
		return []int{}
	}
	leftBounds := make([]LeftBound, n)
	for i := 0; i < n; i++ {
		leftBounds[i] = LeftBound{intervals[i][0], i}
	}

	sort.Sort(Bounds(leftBounds))

	res := make([]int, n)
	for i := 0; i < n; i++ {
		idx, found := halfFind(leftBounds, intervals[i][1])
		if found {
			res[i] = leftBounds[idx].Index
		} else {
			if idx+1 >= n {
				res[i] = -1
			} else {
				res[i] = leftBounds[idx+1].Index
			}
		}
	}
	return res
}

func main() {
	intervals := [][]int{{3, 4}, {2, 3}, {1, 2}}
	fmt.Println(findRightInterval(intervals))

	intervals = [][]int{{1, 4}, {2, 3}, {3, 4}}
	fmt.Println(findRightInterval(intervals))
}
