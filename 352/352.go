package main

import (
	"fmt"
)

type SummaryRanges struct {
	Intervals [][]int
}

/** Initialize your data structure here. */
func Constructor() SummaryRanges {
	return SummaryRanges{make([][]int, 0)}
}

func addAfterIdx(val int, intervals [][]int, idx int) [][]int {
	if idx == -1 {
		if intervals[0][0]-1 > val {
			return append([][]int{{val, val}}, intervals...)
		} else {
			intervals[0][0] = val
			return intervals
		}
	}
	if idx+1 == len(intervals) {
		if intervals[idx][1]+1 < val {
			return append(intervals, []int{val, val})
		} else {
			intervals[idx][1] = val
			return intervals
		}
	}
	if intervals[idx][1]+1 < val {
		if val < intervals[idx+1][0]-1 {
			temp := append([][]int{{val, val}}, intervals[idx+1:]...)
			intervals = append(intervals[:idx+1], temp...)
		} else {
			intervals[idx+1][0] = val
		}
	} else {
		intervals[idx][1] = val
		if val == intervals[idx+1][0]-1 {
			intervals[idx][1] = intervals[idx+1][1]
			intervals = append(intervals[:idx+1], intervals[idx+2:]...)
		}
	}
	return intervals
}

func (this *SummaryRanges) AddNum(val int) {
	if len(this.Intervals) == 0 {
		this.Intervals = append(this.Intervals, []int{val, val})
		return
	}

	left := 0
	right := len(this.Intervals) - 1

	for right-left+1 > 2 {
		mid := (left + right) / 2
		if this.Intervals[mid][0] > val {
			right = mid - 1
		} else if this.Intervals[mid][1] < val {
			left = mid + 1
		} else {
			return
		}
	}

	if right-left == 0 {
		if this.Intervals[left][0] > val {
			this.Intervals = addAfterIdx(val, this.Intervals, left-1)
		} else if this.Intervals[left][1] < val {
			this.Intervals = addAfterIdx(val, this.Intervals, left)
		}
		return
	}

	if right-left == 1 {
		if this.Intervals[left][0] > val {
			this.Intervals = addAfterIdx(val, this.Intervals, left-1)
		} else if this.Intervals[left][1] < val && this.Intervals[right][0] > val {
			this.Intervals = addAfterIdx(val, this.Intervals, left)
		} else if this.Intervals[right][1] < val {
			this.Intervals = addAfterIdx(val, this.Intervals, right)
		}
		return
	}
}

func (this *SummaryRanges) GetIntervals() [][]int {
	return this.Intervals
}

/**
 * Your SummaryRanges object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(val);
 * param_2 := obj.GetIntervals();
 */

func main() {
	obj := Constructor()
	obj.AddNum(1)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(3)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(7)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(2)
	fmt.Println(obj.GetIntervals())
	obj.AddNum(6)
	fmt.Println(obj.GetIntervals())
}
