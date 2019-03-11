package main

import (
	"fmt"
)

func minSubArrayLen(s int, nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	minLen := numsLen + 1
	left := 0
	right := 0
	sum := nums[0]
	for right < numsLen {
		if sum >= s {
			if minLen > (right - left + 1) {
				minLen = right - left + 1
			}
			left++
			if right < left {
				right = left
				if left < numsLen {
					sum = nums[left]
				} else {
					sum = 0
				}
			} else {
				sum -= nums[left-1]
			}
		} else {
			right++
			if right < numsLen {
				sum += nums[right]
			}
		}
	}
	if minLen > numsLen {
		minLen = 0
	}
	return minLen
}

func main() {
	nums := []int{2, 3, 1, 2, 4, 3}
	s := 100
	fmt.Println(minSubArrayLen(s, nums))
}
