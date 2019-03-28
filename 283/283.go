package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	start := 0
	for i := range nums {
		if nums[i] != 0 {
			if i > start {
				nums[start], nums[i] = nums[i], nums[start]
			}
			start++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
