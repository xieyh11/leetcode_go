package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	stepOne := nums[0]
	stepTwo := nums[0]
	for true {
		stepOne = nums[stepOne]
		stepTwo = nums[nums[stepTwo]]
		if stepOne == stepTwo {
			break
		}
	}
	stepOne = nums[0]
	for stepOne != stepTwo {
		stepOne = nums[stepOne]
		stepTwo = nums[stepTwo]
	}
	return stepOne
}

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}
