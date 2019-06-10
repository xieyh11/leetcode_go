package main

import (
	"fmt"
	"sort"
)

type Interval [][]int

func (a Interval) Len() int {
	return len(a)
}

func (a Interval) Less(i, j int) bool {
	return a[i][1] < a[j][1]
}

func (a Interval) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Sort(Interval(intervals))

	cover := 0
	first := false
	res := 0
	for start := 0; start < len(intervals); start++ {
		if !first {
			cover = intervals[start][1]
			first = true
		} else if cover <= intervals[start][0] {
			cover = intervals[start][1]
		} else {
			res++
		}
	}
	return res
}

func main() {
	intervals := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}
	fmt.Println(eraseOverlapIntervals(intervals))
}
