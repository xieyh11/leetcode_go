package main

import (
	"fmt"
	"sort"
)

func twoSum(nums []int, numOfNums map[int]int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i += numOfNums[nums[i]] {
		tempj := target - nums[i]
		if tempj > nums[len(nums)-1] {
			continue
		} else if tempj < nums[i] {
			break
		} else if tempj == nums[i] {
			if numOfNums[nums[i]] > 1 {
				tempRes := make([]int, 2)
				tempRes[0], tempRes[1] = tempj, tempj
				res = append(res, tempRes)
				break
			}
		} else {
			if num, ok := numOfNums[tempj]; ok && num > 0 {
				tempRes := make([]int, 2)
				tempRes[0], tempRes[1] = nums[i], tempj
				res = append(res, tempRes)
			}
		}
	}
	return res
}

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	split := float64(target) / 4

	res := make([][]int, 0)

	numOfNums := make(map[int]int)
	for _, num := range nums {
		if _, ok := numOfNums[num]; ok {
			numOfNums[num]++
		} else {
			numOfNums[num] = 1
		}
	}

	for left := 0; left < len(nums) && float64(nums[left]) <= split; left += numOfNums[nums[left]] {
		numOfNums[nums[left]]--

		for right := len(nums) - 1; right > left && (float64(nums[right]) > split || (nums[right] == nums[left] && numOfNums[nums[right]] > 0)); right -= numOfNums[nums[right]] {
			numOfNums[nums[right]]--
			tempRes := twoSum(nums[left+1:right], numOfNums, target-nums[left]-nums[right])
			for idx, oneRes := range tempRes {
				oneRes = append(oneRes, nums[right])
				oneRes = append([]int{nums[left]}, oneRes...)
				tempRes[idx] = oneRes
			}
			res = append(res, tempRes...)
			numOfNums[nums[right]]++
		}

		numOfNums[nums[left]]++
	}
	return res
}

func main() {
	nums := []int{4, -3, -8, 5, 2, 10, 6, 3, -4, -4, 4, 3, 0, -10, 0, -6, -5, -3}
	target := -19

	fmt.Println(fourSum(nums, target))
}
