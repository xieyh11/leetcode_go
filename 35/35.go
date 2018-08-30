package main

import (
	"fmt"
)

func halfSearch(nums []int, target int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0, true
		} else {
			return 0, false
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0, true
		}
		if nums[1] == target {
			return 1, true
		}
		if target < nums[0] {
			return 0, false
		}
		return 1, false
	}
	middleIdx := len(nums) / 2
	if nums[middleIdx] == target {
		return middleIdx, true
	} else if nums[middleIdx] < target {
		if temp, ok := halfSearch(nums[middleIdx+1:], target); temp != -1 {
			return temp + middleIdx + 1, ok
		} else {
			return temp, ok
		}
	} else {
		return halfSearch(nums[:middleIdx], target)
	}
}

func searchInsert(nums []int, target int) int {
	res, ok := halfSearch(nums, target)
	if ok {
		return res
	} else {
		if nums[res] > target {
			return res
		} else {
			return res + 1
		}
	}
}

func main() {
	nums := []int{1, 3, 5, 6}
	target := 0

	fmt.Println(searchInsert(nums, target))
}
