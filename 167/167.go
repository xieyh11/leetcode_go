package main

import (
	"fmt"
)

func halfSearch(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		}
		if nums[1] == target {
			return 1
		}
		return -1
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	} else if target < nums[mid] {
		return halfSearch(nums[:mid], target)
	} else {
		temp := halfSearch(nums[mid+1:], target)
		if temp != -1 {
			return mid + 1 + temp
		}
	}
	return -1
}

func twoSum(numbers []int, target int) []int {
	for i := len(numbers) - 1; i > 0; i-- {
		remain := target - numbers[i]
		if remain <= numbers[i] && remain >= numbers[0] {
			res := halfSearch(numbers[:i], remain)
			if res != -1 {
				return []int{res + 1, i + 1}
			}
		}
	}
	return []int{}
}

func main() {
	nums := []int{12, 13, 23, 28, 43, 44, 59, 60, 61, 68, 70, 86, 88, 92, 124, 125, 136, 168, 173, 173, 180, 199, 212, 221, 227, 230, 277, 282, 306, 314, 316, 321, 325, 328, 336, 337, 363, 365, 368, 370, 370, 371, 375, 384, 387, 394, 400, 404, 414, 422, 422, 427, 430, 435, 457, 493, 506, 527, 531, 538, 541, 546, 568, 583, 585, 587, 650, 652, 677, 691, 730, 737, 740, 751, 755, 764, 778, 783, 785, 789, 794, 803, 809, 815, 847, 858, 863, 863, 874, 887, 896, 916, 920, 926, 927, 930, 933, 957, 981, 997}
	target := 542
	fmt.Println(twoSum(nums, target))
}
