package main

import (
	"fmt"
)

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	res[0] = 1
	countZero := 0
	zeroIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			countZero++
			zeroIdx = i
		} else {
			res[0] *= nums[i]
		}
	}
	if nums[0] == 0 {
		countZero++
		zeroIdx = 0
	}
	if countZero == 0 {
		for i := 1; i < len(nums); i++ {
			res[i] = res[i-1] / nums[i] * nums[i-1]
		}
	} else if countZero == 1 {
		if zeroIdx != 0 {
			res[zeroIdx] = res[0] * nums[0]
			res[0] = 0
		}
	} else {
		res[0] = 0
	}
	return res
}

func main() {
	nums := []int{0, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
