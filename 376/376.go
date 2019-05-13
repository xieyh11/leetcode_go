package main

import (
	"fmt"
)

func ascendFind(nums [][2]int, num int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}
	if len(nums) == 1 {
		if num < nums[0][0] {
			return -1, false
		}
		return 0, num == nums[0][0]
	}
	if len(nums) == 2 {
		if num < nums[0][0] {
			return -1, false
		} else {
			if num < nums[1][0] {
				return 0, num == nums[0][0]
			}
			return 1, num == nums[1][0]
		}
	}

	mid := len(nums) / 2
	if nums[mid][0] == num {
		return mid, true
	} else if nums[mid][0] > num {
		return ascendFind(nums[:mid], num)
	} else {
		temp, tempFound := ascendFind(nums[mid+1:], num)
		return temp + mid + 1, tempFound
	}
}

func descendFind(nums [][2]int, num int) (int, bool) {
	if len(nums) == 0 {
		return -1, false
	}
	if len(nums) == 1 {
		if num > nums[0][0] {
			return -1, false
		}
		return 0, num == nums[0][0]
	}
	if len(nums) == 2 {
		if num > nums[0][0] {
			return -1, false
		} else {
			if num > nums[1][0] {
				return 0, num == nums[0][0]
			}
			return 1, num == nums[1][0]
		}
	}

	mid := len(nums) / 2
	if nums[mid][0] == num {
		return mid, true
	} else if nums[mid][0] < num {
		return descendFind(nums[:mid], num)
	} else {
		temp, tempFound := descendFind(nums[mid+1:], num)
		return temp + mid + 1, tempFound
	}
}

func wiggleMaxLength(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	raiseLast := [][2]int{{nums[0], 1}}
	decreaseLast := [][2]int{{nums[0], 1}}

	for i := 1; i < len(nums); i++ {
		raiseIdx, raiseFound := ascendFind(raiseLast, nums[i])
		last := raiseLast[len(raiseLast)-1]
		last[0] = nums[i]

		if raiseFound {
			raiseLast[raiseIdx] = last
			raiseLast = raiseLast[:raiseIdx+1]
			raiseIdx--
		} else if raiseIdx < len(raiseLast)-1 {
			raiseLast[raiseIdx+1] = last
			raiseLast = raiseLast[:raiseIdx+2]
		}

		decreaseIdx, decreaseFound := descendFind(decreaseLast, nums[i])
		last = decreaseLast[len(decreaseLast)-1]
		last[0] = nums[i]

		if decreaseFound {
			decreaseLast[decreaseIdx] = last
			decreaseLast = decreaseLast[:decreaseIdx+1]
			decreaseIdx--
		} else if decreaseIdx < len(decreaseLast)-1 {
			decreaseLast[decreaseIdx+1] = last
			decreaseLast = decreaseLast[:decreaseIdx+2]
		}

		newRaise := [2]int{nums[i], 1}
		if decreaseIdx >= 0 {
			newRaise[1] += decreaseLast[decreaseIdx][1]
		}
		newDecrease := [2]int{nums[i], 1}
		if raiseIdx >= 0 {
			newDecrease[1] += raiseLast[raiseIdx][1]
		}
		if raiseIdx < len(raiseLast)-1 {
			if newRaise[1] > raiseLast[raiseIdx+1][1] {
				raiseLast[raiseIdx+1] = newRaise
			}
		} else if newRaise[1] > raiseLast[raiseIdx][1] {
			raiseLast = append(raiseLast, newRaise)
		}
		if decreaseIdx < len(decreaseLast)-1 {
			if newDecrease[1] > decreaseLast[decreaseIdx+1][1] {
				decreaseLast[decreaseIdx+1] = newDecrease
			}
		} else if newDecrease[1] > decreaseLast[decreaseIdx][1] {
			decreaseLast = append(decreaseLast, newDecrease)
		}
	}
	if raiseLast[len(raiseLast)-1][1] > decreaseLast[len(decreaseLast)-1][1] {
		return raiseLast[len(raiseLast)-1][1]
	}
	return decreaseLast[len(decreaseLast)-1][1]
}

func main() {
	nums := []int{1, 7, 4, 9, 2, 5}
	fmt.Println(wiggleMaxLength(nums))

	nums = []int{1, 17, 5, 10, 13, 15, 10, 5, 16, 8}
	fmt.Println(wiggleMaxLength(nums))

	nums = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(wiggleMaxLength(nums))
}
