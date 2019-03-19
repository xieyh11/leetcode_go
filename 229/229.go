package main

import (
	"fmt"
)

func findFirst(nums []int) int {
	count := 0
	majority := nums[0]
	for _, num := range nums {
		if count == 0 {
			majority = num
		}
		if num == majority {
			count += 2
		} else {
			count--
		}
	}
	return majority
}

func findSecond(nums []int, except int) (int, bool) {
	count := 0
	exceptTimes := 0
	majority := nums[0]
	for _, num := range nums {
		if num != except {
			if count == 0 {
				majority = num
			}
			if num == majority {
				count++
			} else {
				count--
			}
		} else {
			exceptTimes++
		}
	}
	return majority, exceptTimes > (len(nums) / 3)
}

func isMajority(nums []int, match int) bool {
	times := 0
	for _, num := range nums {
		if num == match {
			times++
		}
	}
	return times > (len(nums) / 3)
}

func majorityElement(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	first := findFirst(nums)
	second, firstIs := findSecond(nums, first)
	secondIs := isMajority(nums, second)
	if firstIs && secondIs {
		if first != second {
			return []int{first, second}
		} else {
			return []int{first}
		}
	} else if firstIs {
		return []int{first}
	} else if secondIs {
		return []int{second}
	}
	return []int{}
}

func main() {
	nums := []int{3, 3, 2}
	fmt.Println(majorityElement(nums))
}
