package main

import (
	"fmt"
	"strconv"
	"strings"
)

func compareSameLen(s1, s2 string) int {
	for i := range s1 {
		if s1[i] > s2[i] {
			return 1
		}
		if s1[i] < s2[i] {
			return -1
		}
	}
	return 0
}

func compareTwo(s1, s2 string) int {
	return compareSameLen(s1+s2, s2+s1)
}

func sort(nums []string) []string {
	if len(nums) <= 1 {
		return nums
	}
	if len(nums) == 2 {
		if compareTwo(nums[0], nums[1]) < 0 {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return nums
	}
	mid := len(nums) / 2
	left := sort(nums[:mid])
	right := sort(nums[mid:])
	res := make([]string, len(nums))
	i, j := 0, 0
	k := 0
	for i < len(left) && j < len(right) {
		if compareTwo(left[i], right[j]) >= 0 {
			res[k] = left[i]
			i++
		} else {
			res[k] = right[j]
			j++
		}
		k++
	}
	for i < len(left) {
		res[k] = left[i]
		i++
		k++
	}
	for j < len(right) {
		res[k] = right[j]
		k++
		j++
	}
	return res
}

func largestNumber(nums []int) string {
	numStr := make([]string, len(nums))
	for i := range nums {
		numStr[i] = strconv.FormatInt(int64(nums[i]), 10)
	}
	numStr = sort(numStr)
	res := strings.Join(numStr, "")
	i := 0
	for i = range res {
		if res[i] != '0' {
			break
		}
	}
	return res[i:]
}

func main() {
	nums := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(nums))
}
