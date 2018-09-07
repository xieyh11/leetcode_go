package main

import (
	"fmt"
)

func sortColors(nums []int) {
	left := 0
	right := len(nums) - 1
	for ; left < len(nums) && nums[left] == 0; left++ {
	}
	for ; right >= 0 && nums[right] == 2; right-- {

	}
	if left < right {
		for i := left + 1; left < right && i < right; {
			switch nums[i] {
			case 0:
				nums[left], nums[i] = nums[i], nums[left]
				for ; left < right && nums[left] == 0; left++ {
				}
				if i <= left {
					i = left + 1
				}
			case 2:
				nums[right], nums[i] = nums[i], nums[right]
				for ; left < right && nums[right] == 2; right-- {
				}
				if i >= right {
					i = right - 1
				}
			default:
				i++
			}
		}
		if left < right {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
}

func main() {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	fmt.Println(nums)
}
