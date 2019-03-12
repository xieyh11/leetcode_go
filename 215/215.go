package main

import (
	"fmt"
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	randIdx := rand.Intn(len(nums))
	less := make([]int, 0)
	more := make([]int, 0)
	equal := 0
	randNum := nums[randIdx]
	for _, num := range nums {
		if num < randNum {
			less = append(less, num)
		} else if num == randNum {
			equal++
		} else {
			more = append(more, num)
		}
	}
	if len(more) >= k {
		return findKthLargest(more, k)
	} else if len(more)+equal >= k {
		return randNum
	} else {
		return findKthLargest(less, k-len(more)-equal)
	}
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	fmt.Println(findKthLargest(nums, k))
}
