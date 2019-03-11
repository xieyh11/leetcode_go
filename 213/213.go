package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func robLine(nums []int) int {
	numLen := len(nums)
	if numLen == 0 {
		return 0
	}
	if numLen == 1 {
		return nums[0]
	}
	if numLen == 2 {
		return getMax(nums[0], nums[1])
	}

	res := make([]int, numLen)
	res[numLen-1] = nums[numLen-1]
	res[numLen-2] = getMax(nums[numLen-2], nums[numLen-1])
	for i := numLen - 3; i >= 0; i-- {
		res[i] = getMax(res[i+1], nums[i]+res[i+2])
	}
	return res[0]
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return getMax(nums[0], nums[1])
	}
	return getMax(robLine(nums[1:]), nums[0]+robLine(nums[2:len(nums)-1]))
}

func main() {
	nums := []int{1, 2, 3, 1}
	fmt.Println(rob(nums))
}
