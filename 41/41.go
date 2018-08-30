package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}
	for i := 0; i < len(nums); {
		num := nums[i]
		if num <= len(nums) && num > 0 {
			if nums[num-1] != num {
				nums[num-1], nums[i] = num, nums[num-1]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	nums := []int{7, 8, 9, 11, 12}
	fmt.Println(firstMissingPositive(nums))
}
