package main

import (
	"fmt"
	"sort"
)

func replica(num, count int) []int {
	res := make([]int, count)
	for i := range res {
		res[i] = num
	}
	return res
}

func combination(num, count int) [][]int {
	res := make([][]int, count+1)
	for i := 0; i <= count; i++ {
		res[i] = replica(num, i)
	}
	return res
}

func recurSol(nums []int, numOf map[int]int) [][]int {
	if len(nums) == 1 {
		return [][]int{{}, {nums[0]}}
	}
	firNum, firCount := nums[0], numOf[nums[0]]
	if len(nums) == firCount {
		return combination(firNum, firCount)
	}
	res := recurSol(nums[firCount:], numOf)
	tempR := make([][]int, 0)
	combFir := combination(firNum, firCount)
	for i := range combFir {
		if len(combFir[i]) != 0 {
			for j := range res {
				tempOne := make([]int, 0)
				tempOne = append(tempOne, combFir[i]...)
				tempOne = append(tempOne, res[j]...)
				tempR = append(tempR, tempOne)
			}
		}
	}
	res = append(res, tempR...)
	return res
}

func subsetsWithDup(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	sort.Ints(nums)
	numOf := make(map[int]int)
	for _, num := range nums {
		numOf[num]++
	}
	return recurSol(nums, numOf)
}

func main() {
	nums := []int{1, 2, 2}
	fmt.Println(subsetsWithDup(nums))
}
