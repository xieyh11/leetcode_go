package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	cur := 0
	for i := 0; i < len(nums); {
		count := 0
		j := i
		for ; j < len(nums) && nums[j] == nums[i]; j++ {
			count++
		}
		nums[cur] = nums[i]
		cur++
		if count >= 2 {
			nums[cur] = nums[i+1]
			cur++
		}
		i = j
	}
	return cur
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
