package main

import (
	"fmt"
	"sort"
)

func combinationSum4(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if target%nums[0] == 0 {
			return 1
		}
		return 0
	}
	sort.Ints(nums)
	res := make([]int, target+1)
	res[0] = 1
	numsIdx := 0
	for i := 1; i <= target; i++ {
		for numsIdx < len(nums) && nums[numsIdx] <= i {
			numsIdx++
		}
		for j := 0; j < numsIdx; j++ {
			res[i] += res[i-nums[j]]
		}
	}
	return res[target]
}

func main() {
	nums := []int{3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25}
	target := 1
	fmt.Println(combinationSum4(nums, target))
}
