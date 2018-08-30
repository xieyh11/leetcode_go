package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{[]int{nums[0]}}
	}
	if len(nums) == 2 {
		return [][]int{[]int{nums[0], nums[1]}, []int{nums[1], nums[0]}}
	}

	res := make([][]int, 0)
	for i := range nums {
		nums[i], nums[0] = nums[0], nums[i]
		tempRes := permute(nums[1:])
		for k, v := range tempRes {
			tempRes[k] = append([]int{nums[0]}, v...)
		}
		res = append(res, tempRes...)
		nums[i], nums[0] = nums[0], nums[i]
	}
	return res
}

func main() {
	nums := []int{5, 4, 6, 2}
	fmt.Println(permute(nums))
}
