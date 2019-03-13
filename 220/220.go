package main

import (
	"fmt"
)

func absoluteDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i, num := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if absoluteDiff(num, nums[j]) <= t {
				return true
			}
		}
	}
	return false
}

func main() {
	nums := []int{1, 5, 9, 1, 5, 9}
	k := 2
	t := 3
	fmt.Println(containsNearbyAlmostDuplicate(nums, k, t))
}
