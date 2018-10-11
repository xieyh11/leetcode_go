package main

import (
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[len(nums1)-m:], nums1[:m])
	i, j := len(nums1)-m, 0
	start := 0
	for i < len(nums1) && j < n {
		if nums1[i] > nums2[j] {
			nums1[start] = nums2[j]
			j++
		} else {
			nums1[start] = nums1[i]
			i++
		}
		start++
	}
	if i < len(nums1) {
		copy(nums1[start:], nums1[i:])
	}
	if j < n {
		copy(nums1[start:], nums2[j:n])
	}
}

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
