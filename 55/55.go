package main

import (
	"fmt"
)

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	covers := make([]bool, len(nums))
	covers[0] = true
	rightmost := 1
	for i := 0; i < len(nums)-1 && rightmost < len(nums); i++ {
		if covers[i] {
			for j := rightmost; j < len(nums) && j-i <= nums[i]; j++ {
				covers[j] = true
				rightmost++
			}
		} else {
			break
		}
	}
	return covers[len(covers)-1]
}

func main() {
	nums := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums))
}
