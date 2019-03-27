package main

import (
	"fmt"
)

func xorAll(nums []int) int {
	res := 0
	for _, num := range nums {
		res ^= num
	}
	return res
}

func singleNumber(nums []int) []int {
	if len(nums) < 2 {
		return []int{}
	}

	remain := xorAll(nums)
	justOne := remain & (remain ^ (remain - 1))
	left := 0
	right := len(nums) - 1
	for left < right {
		if nums[left]&justOne == 0 {
			left++
		} else if nums[right]&justOne != 0 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	if nums[left]&justOne == 0 {
		left++
	}
	return []int{xorAll(nums[:left]), xorAll(nums[left:])}
}

func main() {
	nums := []int{1, 2, 1, 3, 2, 5}
	fmt.Println(singleNumber(nums))
}
