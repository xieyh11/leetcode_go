package main

import (
	"fmt"
)

func rotate(nums []int, k int) {
	k %= len(nums)
	if k == 0 {
		return
	}
	numsLen := len(nums)
	for i, j := 0, numsLen-k-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := numsLen-k, numsLen-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, numsLen-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {
	for k := 1; k < 9; k++ {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		rotate(nums, k)
		fmt.Println(nums)
	}
}
