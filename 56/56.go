package main

import (
	"../common"
	"fmt"
	"sort"
)

type Interval = common.Interval
type IntervalArray []Interval

func (a IntervalArray) Len() int           { return len(a) }
func (a IntervalArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntervalArray) Less(i, j int) bool { return a[i].Start < a[j].Start }

func merge(intervals []Interval) []Interval {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Sort(IntervalArray(intervals))
	res := make([]Interval, 0)
	res = append(res, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start <= res[len(res)-1].End {
			if res[len(res)-1].End < intervals[i].End {
				res[len(res)-1].End = intervals[i].End
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

func main() {
	intervals := []Interval{{1, 2}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(merge(intervals))
}
