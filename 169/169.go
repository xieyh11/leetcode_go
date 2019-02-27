package main

import (
	"fmt"
)

func majoritySol2(nums []int) int {
	count := 0
	candidate := 0
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if num == candidate {
			count++
		} else {
			count--
		}
	}
	return candidate
}

func majorityElement(nums []int) int {
	majority := len(nums) / 2
	count := make(map[int]int, majority+1)
	for _, num := range nums {
		count[num]++
		if count[num] > majority {
			return num
		}
	}
	return 0
}

func main() {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	fmt.Println(majoritySol2(nums))
}
