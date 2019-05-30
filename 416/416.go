package main

import (
	"fmt"
)

func exactSum(nums []int, target int, store map[[2]int]bool) bool {
	if len(nums) == 0 {
		return target == 0
	}
	if len(nums) == 1 {
		return target == nums[0]
	}
	loc := [2]int{len(nums), target}
	if v, ok := store[loc]; ok {
		return v
	}
	res := exactSum(nums[1:], target-nums[0], store) || exactSum(nums[1:], target, store)
	store[loc] = res
	return res
}

func canPartition(nums []int) bool {
	if len(nums) == 0 {
		return true
	}
	if len(nums) == 1 {
		return false
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}
	half := sum >> 1
	store := make(map[[2]int]bool)
	return exactSum(nums, half, store)
}

func main() {
	nums := []int{1, 5, 11, 5}
	fmt.Println(canPartition(nums))

	nums = []int{1, 2, 3, 5}
	fmt.Println(canPartition(nums))

	nums = []int{1, 2, 3, 5, 7}
	fmt.Println(canPartition(nums))
}
