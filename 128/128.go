package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numIndex := make(map[int]int)
	used := make([]bool, len(nums))
	for i := range nums {
		if _, ok := numIndex[nums[i]]; ok {
			used[i] = true
		} else {
			numIndex[nums[i]] = i
		}
	}

	maxLen := 1
	for i := range nums {
		if !used[i] {
			used[i] = true
			consecLen := 1
			for j := nums[i] - 1; true; j-- {
				if idx, ok := numIndex[j]; ok {
					consecLen++
					used[idx] = true
				} else {
					break
				}
			}
			for j := nums[i] + 1; true; j++ {
				if idx, ok := numIndex[j]; ok {
					consecLen++
					used[idx] = true
				} else {
					break
				}
			}
			if maxLen < consecLen {
				maxLen = consecLen
			}
		}
	}
	return maxLen
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
