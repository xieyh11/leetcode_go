package main

import (
	"fmt"
)

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left] == val {
			for right > left && nums[right] == val {
				right--
			}
			if right > left {
				nums[left], nums[right] = nums[right], nums[left]
			} else {
				break
			}
		}
		left++
	}
	if nums[left] == val {
		return left
	} else {
		return left + 1
	}
}

func main() {
	nums := []int{3, 3}
	val := 3
	len := removeElement(nums, val)
	fmt.Println(nums[:len])
}
