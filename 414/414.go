package main

import (
	"fmt"
)

func thirdMax(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	top3 := [3]int{nums[0]}
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] < top3[count-1] {
			if count < 3 {
				top3[count] = nums[i]
				count++
			}
		} else if nums[i] > top3[count-1] {
			if count == 1 {
				top3[count], top3[count-1] = top3[count-1], nums[i]
				count++
			} else if nums[i] < top3[count-2] {
				if count < 3 {
					top3[count], top3[count-1] = top3[count-1], nums[i]
					count++
				} else {
					top3[count-1] = nums[i]
				}
			} else if nums[i] > top3[count-2] {
				if count == 2 {
					top3[count], top3[count-1], top3[count-2] = top3[count-1], top3[count-2], nums[i]
					count++
				} else if nums[i] < top3[count-3] {
					top3[count-1], top3[count-2] = top3[count-2], nums[i]
				} else if nums[i] > top3[count-3] {
					top3[count-1], top3[count-2], top3[count-3] = top3[count-2], top3[count-3], nums[i]
				}
			}
		}
	}
	if count < 3 {
		return top3[0]
	} else {
		return top3[2]
	}
}

func main() {
	nums := []int{3, 2, 1}
	fmt.Println(thirdMax(nums))

	nums = []int{1, 2}
	fmt.Println(thirdMax(nums))

	nums = []int{2, 2, 3, 1}
	fmt.Println(thirdMax(nums))
}
