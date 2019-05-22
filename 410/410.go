package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func recSol(nums []int, m int, sum int, store map[[2]int]int) int {
	if m == 1 {
		return sum
	}
	loc := [2]int{len(nums), m}
	if v, ok := store[loc]; ok {
		return v
	}
	if len(nums) == m {
		max := 0
		for _, num := range nums {
			max = getMax(max, num)
		}
		store[loc] = max
		return max
	}
	each := sum / m
	current := 0
	currentIdx := 0
	for currentIdx < len(nums) && current < each {
		current += nums[currentIdx]
		currentIdx++
	}
	currentIdx--
	current -= nums[currentIdx]
	res := 0
	initial := true
	for currentIdx < len(nums) {
		tempR := recSol(nums[currentIdx:], m-1, sum-current, store)
		if initial {
			initial = false
			res = getMax(current, tempR)
		} else {
			res = getMin(res, getMax(current, tempR))
		}
		if current >= tempR || len(nums)-currentIdx == m-1 {
			break
		}
		current += nums[currentIdx]
		currentIdx++
	}
	store[loc] = res
	return res
}

func splitArray(nums []int, m int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	store := make(map[[2]int]int)
	return recSol(nums, m, sum, store)
}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6}
	m := 20
	fmt.Println(splitArray(nums, m))
}
