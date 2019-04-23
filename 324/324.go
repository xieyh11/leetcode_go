package main

import (
	"fmt"
	"math/rand"
)

func findKth(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		smaller, larger := nums[0], nums[1]
		if smaller > larger {
			smaller, larger = larger, smaller
		}
		if k == 1 {
			return smaller
		}
		return larger
	}
	idx := rand.Intn(len(nums))
	mid := nums[idx]
	former := 0
	latter := len(nums) - 1
	for former < latter {
		if nums[former] <= mid {
			former++
		} else if nums[latter] > mid {
			latter--
		} else {
			nums[former], nums[latter] = nums[latter], nums[former]
			former++
			latter--
		}
	}
	if nums[former] <= mid {
		former++
	}
	lessOrEqual := former
	former, latter = 0, lessOrEqual-1
	for former < latter {
		if nums[former] < mid {
			former++
		} else if nums[latter] == mid {
			latter--
		} else {
			nums[former], nums[latter] = nums[latter], nums[former]
			former++
			latter--
		}
	}
	if nums[former] < mid {
		former++
	}
	less := former
	if k <= less {
		return findKth(nums[:less], k)
	} else if k <= lessOrEqual {
		return mid
	} else {
		return findKth(nums[lessOrEqual:], k-lessOrEqual)
	}
}

func wiggleSort(nums []int) {
	if len(nums) <= 1 {
		return
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}
	mid := findKth(nums, (len(nums)+1)/2)
	res := make([]int, len(nums))
	odd := 0
	even := len(nums) - 1
	if len(nums)%2 != 0 {
		even--
	}
	for _, num := range nums {
		if num < mid {
			res[odd] = num
			odd += 2
		} else if num > mid {
			res[even] = num
			even -= 2
		}
	}
	half := odd
	for odd < len(nums) {
		res[odd] = mid
		odd += 2
	}
	for even >= 0 {
		res[even] = mid
		even -= 2
	}
	if half > 0 && res[half] == res[half-1] {
		copy(nums[len(nums)-half:], res[:half])
		copy(nums[:len(nums)-half], res[half:])
	} else {
		copy(nums, res)
	}
}

func main() {
	nums := []int{4, 5, 5, 5, 5, 6, 6, 6}
	wiggleSort(nums)
	fmt.Println(nums)
}
