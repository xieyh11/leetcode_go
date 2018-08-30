package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if nums[0] == nums[len(nums)-1] {
		if nums[0] == target {
			return []int{0, len(nums) - 1}
		} else {
			return []int{-1, -1}
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return []int{0, 0}
		} else if nums[1] == target {
			return []int{1, 1}
		} else {
			return []int{-1, -1}
		}
	}
	middleIdx := len(nums) / 2
	res := make([]int, 2)
	if nums[middleIdx] == target {
		res1, res2 := searchRange(nums[:middleIdx], target), searchRange(nums[middleIdx+1:], target)
		if res1[0] != -1 {
			res[0] = res1[0]
		} else {
			res[0] = middleIdx
		}
		if res2[1] != -1 {
			res[1] = res2[1] + middleIdx + 1
		} else {
			res[1] = middleIdx
		}
	} else if nums[middleIdx] < target {
		res = searchRange(nums[middleIdx+1:], target)
		if res[0] != -1 {
			res[0] += middleIdx + 1
			res[1] += middleIdx + 1
		}
	} else {
		res = searchRange(nums[:middleIdx], target)
	}
	return res
}

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	fmt.Println(searchRange(nums, target))
}
