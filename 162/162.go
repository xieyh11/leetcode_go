package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := len(nums) - 1

	for left < right-1 {
		mid := (left + right) / 2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		} else if nums[mid] > nums[mid-1] && nums[mid] < nums[mid+1] {
			left = mid
		} else {
			right = mid
		}
	}
	if right == left {
		if left == 0 {
			return 0
		} else {
			return -1
		}
	} else {
		if left == 0 && nums[0] > nums[1] {
			return 0
		}
		if right == len(nums)-1 && nums[right] > nums[right-1] {
			return right
		}
		return -1
	}
}

func main() {
	nums := []int{1, 2, 1, 3, 5, 6, 4}
	fmt.Println(findPeakElement(nums))
}
