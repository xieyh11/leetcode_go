package main

import (
	"fmt"
	"sort"
)

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func divideFind(nums []int, start, end, target int) (int, bool) {
	if start > end {
		if start >= len(nums) {
			return end, false
		}
		if end < 0 {
			return start, false
		}
	}
	if start == end {
		return start, nums[start] == target
	}
	if start == end-1 {
		if nums[start] == target {
			return start, true
		}
		if nums[start] > target {
			return start, false
		}
		if nums[end] == target {
			return end, true
		}
		return end, false
	}
	middle := (start + end) / 2
	if nums[middle] == target {
		return middle, true
	} else if nums[middle] < target {
		return divideFind(nums, middle+1, end, target)
	} else {
		return divideFind(nums, start, middle-1, target)
	}
}

func threeSumClosest(nums []int, target int) int {
	numOfNums := make(map[int]int)
	sort.Ints(nums)
	diff := 0
	for diff < len(nums) {
		start := diff
		num := nums[diff]
		j := diff + 1
		for ; j < len(nums) && nums[j] == num; j++ {
		}
		numOfNums[num] = j - start
		diff = j
	}
	minDiff := MaxInt
	minDiffSum := 0
	for i := 0; i < len(nums); i += numOfNums[nums[i]] {
		numI := nums[i]
		numOfNums[nums[i]]--
		for j := i + 1; j < len(nums)-1; j += numOfNums[nums[j]] {
			numJ := nums[j]
			numR := target - numI - numJ
			if numR < numJ {
				if nums[j+1]-numR < minDiff {
					minDiff = nums[j+1] - numR
					minDiffSum = target + minDiff
				}
				break
			}
			closeR, okR := divideFind(nums, j+1, len(nums)-1, numR)
			if okR {
				minDiff = 0
				minDiffSum = target
				break
			} else {
				if closeR == j+1 && numR < nums[closeR] {
					if nums[closeR]-numR < minDiff {
						minDiff = nums[closeR] - numR
						minDiffSum = target + minDiff
					}
				} else if closeR == len(nums)-1 && numR > nums[closeR] {
					if numR-nums[closeR] < minDiff {
						minDiff = numR - nums[closeR]
						minDiffSum = target - minDiff
					}
				} else {
					var left, right int
					if nums[closeR] > numR {
						right = nums[closeR]
						left = nums[closeR-1]
					} else {
						left = nums[closeR]
						right = nums[closeR+1]
					}
					if numR-left < minDiff {
						minDiff = numR - left
						minDiffSum = target - minDiff
					}
					if right-numR < minDiff {
						minDiff = right - numR
						minDiffSum = target + minDiff
					}
				}
			}
		}
		if minDiff == 0 {
			break
		}
		numOfNums[nums[i]]++
	}
	return minDiffSum
}

func main() {
	nums := []int{-1, 0, 1, 1, 55}
	target := 3
	fmt.Println(threeSumClosest(nums, target))
}
