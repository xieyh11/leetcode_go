package main

import (
	"../common"
	"fmt"
)

type Interval = common.Interval

func insert(intervals []Interval, newInterval Interval) []Interval {
	i := 0
	for ; i < len(intervals); i++ {
		if intervals[i].End >= newInterval.Start {
			break
		}
	}
	if i < len(intervals) {
		if intervals[i].End <= newInterval.End {
			if intervals[i].Start > newInterval.Start {
				intervals[i].Start = newInterval.Start
			}
			intervals[i].End = newInterval.End
			j := i + 1
			for ; j < len(intervals); j++ {
				if intervals[j].Start <= intervals[i].End {
					if intervals[j].End >= intervals[i].End {
						intervals[i].End = intervals[j].End
						j++
						break
					}
				} else {
					break
				}
			}
			i++
			if j < len(intervals) {
				for ; j < len(intervals); j++ {
					intervals[i] = intervals[j]
					i++
				}
			}
			intervals = intervals[:i]
		} else if intervals[i].Start <= newInterval.End {
			if intervals[i].Start > newInterval.Start {
				intervals[i].Start = newInterval.Start
			}
		} else {
			intervals = append(intervals, Interval{})
			for j := len(intervals) - 1; j > i; j-- {
				intervals[j] = intervals[j-1]
			}
			intervals[i] = newInterval
		}
	} else {
		intervals = append(intervals, newInterval)
	}
	return intervals
}

func main() {
	intervals := []Interval{{1, 5}}
	newInterval := Interval{0, 6}

	fmt.Println(insert(intervals, newInterval))
}
