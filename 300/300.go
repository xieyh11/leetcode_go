package main

import (
	"fmt"
)

func halfFind(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] >= target {
			return 0
		}
		return 1
	}
	if len(nums) == 2 {
		if nums[0] >= target {
			return 0
		}
		if nums[1] >= target {
			return 1
		}
		return 2
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return halfFind(nums[:mid], target)
	} else {
		return mid + 1 + halfFind(nums[mid+1:], target)
	}
}

func lengthOfLIS(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	res := make([]int, 1)
	res[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		idx := halfFind(res, nums[i])
		if idx < len(res) {
			res[idx] = nums[i]
		} else {
			res = append(res, nums[i])
		}
	}
	return len(res)
}

func main() {
	nums := []int{10, 9, 2, 5, 3, 7, 101, 18}
	fmt.Println(lengthOfLIS(nums))
}
