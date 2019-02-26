package main

import (
	"fmt"
)

func findMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	leftE := nums[0]
	rightE := nums[len(nums)-1]

	if len(nums) > 1 && leftE < rightE {
		return leftE
	}

	left := 0
	right := len(nums) - 1

	for left < right-1 {
		midIdx := (left + right) / 2
		if nums[midIdx] < nums[midIdx-1] && nums[midIdx] < nums[midIdx+1] {
			return nums[midIdx]
		} else if nums[midIdx] <= rightE {
			right = midIdx - 1
		} else {
			left = midIdx + 1
		}
	}
	numOf := right - left + 1
	if numOf == 1 {
		return nums[left]
	}
	if numOf == 2 {
		if nums[left] < nums[right] {
			return nums[left]
		} else {
			return nums[right]
		}
	}
	return nums[left]
}

func main() {
	nums := []int{3, 4, 5, 1, 2}
	fmt.Println(findMin(nums))
}
