package main

import (
	"fmt"
)

func searchLess(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
	}
	return -1
}

func halfSearch(nums []int, target int) int {
	if len(nums) <= 2 {
		return searchLess(nums, target)
	}
	middleIdx := len(nums) / 2
	if nums[middleIdx] == target {
		return middleIdx
	} else if nums[middleIdx] < target {
		if temp := halfSearch(nums[middleIdx+1:], target); temp != -1 {
			return temp + middleIdx + 1
		} else {
			return -1
		}
	} else {
		return halfSearch(nums[:middleIdx], target)
	}
}

func search(nums []int, target int) int {
	if len(nums) <= 2 {
		return searchLess(nums, target)
	}

	if nums[0] < nums[len(nums)-1] {
		return halfSearch(nums, target)
	}

	middleIdx := len(nums) / 2
	if nums[middleIdx] == target {
		return middleIdx
	} else if nums[middleIdx] < target {
		if nums[middleIdx] >= nums[0] {
			if temp := search(nums[middleIdx+1:], target); temp != -1 {
				return temp + middleIdx + 1
			} else {
				return -1
			}
		} else {
			if nums[0] == target {
				return 0
			} else if nums[0] < target {
				return search(nums[:middleIdx], target)
			} else {
				if temp := halfSearch(nums[middleIdx+1:], target); temp != -1 {
					return temp + middleIdx + 1
				} else {
					return -1
				}
			}
		}
	} else {
		if nums[middleIdx] <= nums[len(nums)-1] {
			return search(nums[:middleIdx], target)
		} else {
			if nums[len(nums)-1] == target {
				return len(nums) - 1
			} else if nums[len(nums)-1] > target {
				if temp := search(nums[middleIdx+1:], target); temp != -1 {
					return temp + middleIdx + 1
				} else {
					return -1
				}
			} else {
				return halfSearch(nums[:middleIdx], target)
			}
		}
	}
}

func main() {
	nums := []int{7, 8, 1, 2, 3, 4, 5, 6}
	target := 2
	fmt.Println(search(nums, target))
}
