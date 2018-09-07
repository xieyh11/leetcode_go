package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) == 1 {
		return [][]int{nums, {}}
	}
	res := subsets(nums[1:])
	temp := make([][]int, 0)
	for i := range res {
		temp = append(temp, append([]int{nums[0]}, res[i]...))
	}
	res = append(res, temp...)
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(subsets(nums))
}
