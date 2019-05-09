package main

import (
	"fmt"
	"sort"
)

func largestDivisibleSubset(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	sort.Ints(nums)
	length := make([]int, len(nums))
	nextIdx := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if length[i] == 0 {
			length[i] = 1
			nextIdx[i] = -1
		}
		for j := i + 1; j < len(nums); j++ {
			if nums[j]%nums[i] == 0 {
				if length[i] < length[j]+1 {
					length[i] = length[j] + 1
					nextIdx[i] = j
				}
			}
		}
	}
	maxLen := 0
	maxIdx := 0
	for i := range length {
		if maxLen < length[i] {
			maxLen = length[i]
			maxIdx = i
		}
	}

	res := []int{}
	for maxIdx >= 0 {
		res = append(res, nums[maxIdx])
		maxIdx = nextIdx[maxIdx]
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(largestDivisibleSubset(nums))

	nums = []int{1, 2, 4, 8}
	fmt.Println(largestDivisibleSubset(nums))

	nums = []int{1, 2, 4, 5, 8}
	fmt.Println(largestDivisibleSubset(nums))
}
