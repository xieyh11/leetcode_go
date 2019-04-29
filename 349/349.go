package main

import (
	"fmt"
)

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	cache := make(map[int]bool)
	for _, num := range nums1 {
		cache[num] = true
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if v, ok := cache[num]; ok && v {
			cache[num] = false
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums0 := []int{1, 2, 2, 1}
	nums1 := []int{2, 2}
	fmt.Println(intersection(nums0, nums1))
}
