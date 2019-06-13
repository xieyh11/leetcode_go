package main

import (
	"fmt"
)

func findDuplicates(nums []int) []int {
	res := make([]int, 0)
	left := 0
	for left < len(nums) {
		for nums[left] > 0 && nums[left] != left+1 {
			ex := nums[left] - 1
			if nums[ex] > 0 {
				if nums[ex] == nums[left] {
					res = append(res, nums[ex])
					nums[ex] = 0
				} else {
					nums[ex], nums[left] = nums[left], nums[ex]
				}
			} else {
				break
			}
		}
		left++
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDuplicates(nums))
}
