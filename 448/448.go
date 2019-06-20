package main

import (
	"fmt"
)

func findDisappearedNumbers(nums []int) []int {
	left := 0
	n := len(nums)
	for left < n {
		if nums[left] != left+1 {
			num := nums[left]
			if nums[num-1] != num {
				nums[left], nums[num-1] = nums[num-1], num
			} else {
				left++
			}
		} else {
			left++
		}
	}
	res := make([]int, 0)
	for i := range nums {
		if i+1 != nums[i] {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	nums := []int{4, 3, 2, 7, 8, 2, 3, 1}
	fmt.Println(findDisappearedNumbers(nums))
}
