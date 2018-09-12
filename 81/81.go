package main

import (
	"fmt"
)

func halfSearch(nums []int, start, end int, target int) bool {
	length := 0
	if end >= start {
		length = end - start + 1
	} else {
		length = len(nums) - start + end + 1
	}
	if length == 1 {
		return nums[start] == target
	}
	if length == 2 {
		return (nums[start] == target) || (nums[end] == target)
	}

	mid := length / 2
	mid += start
	if mid >= len(nums) {
		mid -= len(nums)
	}
	if nums[mid] == target {
		return true
	} else if nums[mid] > target {
		mid--
		if mid < 0 {
			mid = len(nums) - 1
		}
		return halfSearch(nums, start, mid, target)
	} else {
		mid++
		if mid >= len(nums) {
			mid = 0
		}
		return halfSearch(nums, mid, end, target)
	}
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	if len(nums) == 1 {
		return nums[0] == target
	}
	if len(nums) == 2 {
		return (nums[0] == target) || (nums[1] == target)
	}

	start := 0
	for ; start < len(nums); start++ {
		if start > 0 && nums[start] < nums[start-1] {
			break
		}
	}
	if start < len(nums) {
		return halfSearch(nums, start, start-1, target)
	} else {
		return halfSearch(nums, 0, len(nums)-1, target)
	}
}

func main() {
	nums := []int{2, 5, 6, 0, 0, 1, 2}
	target := 0
	fmt.Println(search(nums, target))
}
