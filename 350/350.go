package main

import (
	"fmt"
)

func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	cache := make(map[int]int)
	for _, num := range nums1 {
		cache[num]++
	}

	res := make([]int, 0)
	for _, num := range nums2 {
		if v, ok := cache[num]; ok && v > 0 {
			cache[num] = v - 1
			res = append(res, num)
		}
	}
	return res
}

func main() {
	nums0 := []int{1, 2, 2, 1}
	nums1 := []int{2, 2}
	fmt.Println(intersect(nums0, nums1))

	nums2 := []int{4, 9, 5}
	nums3 := []int{9, 4, 9, 8, 4}
	fmt.Println(intersect(nums2, nums3))
}
