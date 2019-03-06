package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return getMax(nums[0], nums[1])
	}
	res := make([]int, len(nums))
	res[len(nums)-1] = nums[len(nums)-1]
	res[len(nums)-2] = getMax(res[len(nums)-1], nums[len(nums)-2])
	for i := len(nums) - 3; i >= 0; i-- {
		res[i] = getMax(nums[i]+res[i+2], res[i+1])
	}
	return res[0]
}

func main() {
	nums := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(nums))
}
