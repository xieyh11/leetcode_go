package main

import (
	"fmt"
)

func reverseIntArray(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 1
	for ; i > 0; i-- {
		if nums[i] > nums[i-1] {
			break
		}
	}
	if i == 0 {
		reverseIntArray(nums)
	} else {
		j := i + 1
		for ; j < len(nums); j++ {
			if nums[j] <= nums[i-1] {
				j--
				break
			}
		}
		if j == len(nums) {
			j--
		}
		nums[i-1], nums[j] = nums[j], nums[i-1]
		reverseIntArray(nums[i:])
	}
}

func main() {
	nums := []int{1, 5, 1}
	for i := 0; i < 10; i++ {
		nextPermutation(nums)
		fmt.Println(nums)
	}
}
