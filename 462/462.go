package main

import (
	"fmt"
	"math/rand"
)

func findMedian(nums []int, k int) int {
	idx := rand.Intn(len(nums))
	idxNum := nums[idx]
	left := 0
	right := len(nums) - 1
	for left <= right {
		if nums[left] <= idxNum {
			left++
		} else if nums[right] > idxNum {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	lessEqual := left
	left = 0
	right = lessEqual - 1
	for left <= right {
		if nums[left] < idxNum {
			left++
		} else if nums[right] == idxNum {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
			right--
		}
	}
	less := left
	if less >= k {
		return findMedian(nums[:less], k)
	} else if lessEqual >= k {
		return idxNum
	} else {
		return findMedian(nums[lessEqual:], k-lessEqual)
	}
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func minMoves2(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	mid := findMedian(nums, (len(nums)+1)/2)
	res := 0
	for _, num := range nums {
		res += absDiff(num, mid)
	}
	return res
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(minMoves2(nums))
}
