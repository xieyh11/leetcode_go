package main

import (
	"fmt"
)

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	res := make([]int, len(nums))
	for i := len(nums) - 2; i >= 0; i-- {
		minS := res[i+1] + 1
		for j := i + 1; j < len(nums) && j-i <= nums[i]; j++ {
			if minS > 1+res[j] {
				minS = 1 + res[j]
			}
		}
		res[i] = minS
	}
	return res[0]
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(jump(nums))
}
