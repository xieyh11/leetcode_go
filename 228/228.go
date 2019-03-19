package main

import (
	"fmt"
	"strconv"
)

func summaryRanges(nums []int) []string {
	start := 0
	numsLen := len(nums)
	res := make([]string, 0)
	for start < numsLen {
		end := start + 1
		for end < numsLen && nums[end] == nums[end-1]+1 {
			end++
		}
		if end > start+1 {
			res = append(res, strconv.FormatInt(int64(nums[start]), 10)+"->"+strconv.FormatInt(int64(nums[end-1]), 10))
		} else {
			res = append(res, strconv.FormatInt(int64(nums[start]), 10))
		}
		start = end
	}
	return res
}

func main() {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	fmt.Println(summaryRanges(nums))
}
