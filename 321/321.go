package main

import (
	"fmt"
)

func oneRecSol(nums []int, k int) []int {
	if len(nums) == k {
		return nums
	}
	if k == 0 || k > len(nums) {
		return []int{}
	}
	max := nums[0]
	maxIdx := 0
	for i := 1; i <= len(nums)-k && max < 9; i++ {
		if nums[i] > max {
			max = nums[i]
			maxIdx = i
		}
	}
	res := append([]int{max}, oneRecSol(nums[maxIdx+1:], k-1)...)
	return res
}

func compareArray(nums1, nums2 []int) int {
	less := len(nums1)
	if less > len(nums2) {
		less = len(nums2)
	}
	i := 0
	for ; i < less; i++ {
		if nums1[i] > nums2[i] {
			return 1
		} else if nums1[i] < nums2[i] {
			return -1
		}
	}
	if i == len(nums2) {
		if i == len(nums1) {
			return 0
		}
		return 1
	}
	return -1
}

func mergeToOne(nums1, nums2 []int) []int {
	if len(nums1) == 0 {
		return nums2
	}
	if len(nums2) == 0 {
		return nums1
	}
	res := make([]int, 0)
	for len(nums1) > 0 && len(nums2) > 0 {
		com := compareArray(nums1, nums2)
		if com == 1 || com == 0 {
			res = append(res, nums1[0])
			nums1 = nums1[1:]
		} else {
			res = append(res, nums2[0])
			nums2 = nums2[1:]
		}
	}
	if len(nums1) > 0 {
		res = append(res, nums1...)
	}
	if len(nums2) > 0 {
		res = append(res, nums2...)
	}
	return res
}

func maxArray(a, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			return a
		} else if a[i] < b[i] {
			return b
		}
	}
	return a
}

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	if len(nums1)+len(nums2) < k {
		return []int{}
	}
	res := []int{}
	for i := 0; i <= k && i <= len(nums1); i++ {
		if k-i <= len(nums2) {
			res1 := oneRecSol(nums1, i)
			res2 := oneRecSol(nums2, k-i)
			tempRes := mergeToOne(res1, res2)
			res = maxArray(res, tempRes)
		}
	}
	return res
}

func main() {
	nums1 := []int{3, 3, 3, 2, 3, 7, 3, 8, 6, 0, 5, 0, 7, 8, 9, 2, 9, 6, 6, 9, 9, 7, 9, 7, 6, 1, 7, 2, 7, 5, 5, 1}
	nums2 := []int{5, 6, 4, 9, 6, 9, 2, 2, 7, 5, 4, 3, 0, 0, 1, 7, 1, 8, 1, 5, 2, 5, 7, 0, 4, 3, 8, 7, 3, 8, 5, 3, 8, 3, 4, 0, 2, 3, 8, 2, 7, 1, 2, 3, 8, 7, 6, 7, 1, 1, 3, 9, 0, 5, 2, 8, 2, 8, 7, 5, 0, 8, 0, 7, 2, 8, 5, 6, 5, 9, 5, 1, 5, 2, 6, 2, 4, 9, 9, 7, 6, 5, 7, 9, 2, 8, 8, 3, 5, 9, 5, 1, 8, 8, 4, 6, 6, 3, 8, 4, 6, 6, 1, 3, 4, 1, 6, 7, 0, 8, 0, 3, 3, 1, 8, 2, 2, 4, 5, 7, 3, 7, 7, 4, 3, 7, 3, 0, 7, 3, 0, 9, 7, 6, 0, 3, 0, 3, 1, 5, 1, 4, 5, 2, 7, 6, 2, 4, 2, 9, 5, 5, 9, 8, 4, 2, 3, 6, 1, 9}
	k := 160
	res := maxNumber(nums1, nums2, k)
	expected := []int{9, 9, 9, 9, 9, 9, 9, 7, 8, 7, 6, 1, 7, 2, 7, 5, 5, 1, 5, 2, 5, 7, 1, 0, 4, 3, 8, 7, 3, 8, 5, 3, 8, 3, 4, 0, 2, 3, 8, 2, 7, 1, 2, 3, 8, 7, 6, 7, 1, 1, 3, 9, 0, 5, 2, 8, 2, 8, 7, 5, 0, 8, 0, 7, 2, 8, 5, 6, 5, 9, 5, 1, 5, 2, 6, 2, 4, 9, 9, 7, 6, 5, 7, 9, 2, 8, 8, 3, 5, 9, 5, 1, 8, 8, 4, 6, 6, 3, 8, 4, 6, 6, 1, 3, 4, 1, 6, 7, 0, 8, 0, 3, 3, 1, 8, 2, 2, 4, 5, 7, 3, 7, 7, 4, 3, 7, 3, 0, 7, 3, 0, 9, 7, 6, 0, 3, 0, 3, 1, 5, 1, 4, 5, 2, 7, 6, 2, 4, 2, 9, 5, 5, 9, 8, 4, 2, 3, 6, 1, 9}
	if len(expected) != len(res) {
		fmt.Println("Not Equal")
	}
	for i := range res {
		if res[i] != expected[i] {
			fmt.Println("Not Equal")
			break
		}
	}
	fmt.Println(res)
}
