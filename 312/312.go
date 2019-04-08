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

func maxCoins(nums []int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return 0
	}
	if numsLen == 1 {
		return nums[0]
	}
	if numsLen == 2 {
		first := nums[0]*nums[1] + nums[1]
		second := nums[0]*nums[1] + nums[0]
		return getMax(first, second)
	}

	res := make([][]int, numsLen)
	res[0] = make([]int, numsLen)
	res[0][0] = nums[0] * nums[1]
	for i := 1; i < numsLen-1; i++ {
		res[0][i] = nums[i-1] * nums[i] * nums[i+1]
	}
	res[0][numsLen-1] = nums[numsLen-2] * nums[numsLen-1]

	res[1] = make([]int, numsLen-1)
	res[1][0] = getMax(nums[0]*nums[1]+nums[1]*nums[2], nums[0]*nums[1]*nums[2]+nums[0]*nums[2])
	for i := 1; i < numsLen-2; i++ {
		res[1][i] = getMax(nums[i-1]*nums[i]*nums[i+1]+nums[i-1]*nums[i+1]*nums[i+2], nums[i]*nums[i+1]*nums[i+2]+nums[i-1]*nums[i]*nums[i+2])
	}
	res[1][numsLen-2] = getMax(nums[numsLen-3]*nums[numsLen-2]*nums[numsLen-1]+nums[numsLen-3]*nums[numsLen-1], nums[numsLen-2]*nums[numsLen-1]+nums[numsLen-3]*nums[numsLen-2])

	for k := 2; k < numsLen; k++ {
		res[k] = make([]int, numsLen-k)
		for i := 0; i < numsLen-k; i++ {
			left := 1
			if i > 0 {
				left = nums[i-1]
			}
			right := 1
			if i+k+1 < numsLen {
				right = nums[i+k+1]
			}
			res[k][i] = getMax(left*nums[i]*right+res[k-1][i+1], res[k-1][i]+left*nums[i+k]*right)
			for j := i + 1; j < i+k; j++ {
				res[k][i] = getMax(res[k][i], res[j-i-1][i]+left*nums[j]*right+res[i+k-j-1][j+1])
			}
		}
	}

	return res[numsLen-1][0]
}

func main() {
	nums := []int{3, 1, 5, 8}
	fmt.Println(maxCoins(nums))
}
