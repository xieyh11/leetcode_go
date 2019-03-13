package main

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	numMap := make(map[int]int)
	for i, num := range nums {
		if j, ok := numMap[num]; ok {
			if i-j <= k {
				return true
			}
		}
		numMap[num] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1, 2, 3}
	k := 2
	fmt.Println(containsNearbyDuplicate(nums, k))
}
