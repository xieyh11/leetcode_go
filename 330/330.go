package main

import (
	"fmt"
)

func minPatches(nums []int, n int) int {
	miss := 1
	i := 0
	numsLen := len(nums)
	res := 0
	for miss <= n {
		if i < numsLen && miss >= nums[i] {
			miss += nums[i]
			i++
		} else {
			miss += miss
			res++
		}
	}
	return res
}

func main() {
	nums := []int{}
	n := 6
	fmt.Println(minPatches(nums, n))
}
