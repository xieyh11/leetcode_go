package main

import (
	"fmt"
)

func noZeroSol(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := 1
	for _, num := range nums {
		res *= num
	}
	if res > 0 {
		return res
	}
	left := 1
	for _, num := range nums {
		left *= num
		if left < 0 {
			break
		}
	}
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		right *= nums[i]
		if right < 0 {
			break
		}
	}
	if left > right {
		res /= left
	} else {
		res /= right
	}
	return res
}

func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	formerZero := -1
	if nums[0] == 0 {
		formerZero = 0
	}
	maxRes := nums[0]
	for formerZero < len(nums)-1 {
		temp := formerZero + 1
		for temp < len(nums) && nums[temp] != 0 {
			temp++
		}
		if temp < len(nums) {
			noZeroNums := nums[formerZero+1 : temp]
			if len(noZeroNums) > 0 {
				tempRes := noZeroSol(noZeroNums)
				if tempRes > maxRes {
					maxRes = tempRes
				}
			}
			if maxRes < 0 {
				maxRes = 0
			}
		} else {
			noZeroNums := nums[formerZero+1:]
			if len(noZeroNums) > 0 {
				tempRes := noZeroSol(noZeroNums)
				if tempRes > maxRes {
					maxRes = tempRes
				}
			}
			if formerZero >= 0 && maxRes < 0 {
				maxRes = 0
			}
		}

		formerZero = temp
	}
	return maxRes
}

func main() {
	nums := []int{-2, 0, -1}
	fmt.Println(maxProduct(nums))
}
