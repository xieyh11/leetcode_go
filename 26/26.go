package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	copyAt := 0
	nowAt := 0
	for nowAt < len(nums) {
		numAt := nums[nowAt]
		j := nowAt + 1
		for ; j < len(nums); j++ {
			if numAt != nums[j] {
				break
			}
		}
		if nowAt != copyAt {
			nums[copyAt] = numAt
		}
		copyAt++
		nowAt = j
	}
	return copyAt
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	len := removeDuplicates(nums)
	fmt.Println(nums[:len])
}
