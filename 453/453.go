package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	n := len(nums)
	for i := 1; i < n; i++ {
		min = getMin(min, nums[i])
	}
	res := 0
	for i := range nums {
		res += nums[i] - min
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves(nums))

	nums = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(minMoves(nums))
}
