package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	numsMap := make(map[int]int)
	single := make([]int, 0)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		numOf, ok := numsMap[nums[i]]
		if ok {
			numsMap[nums[i]] = numOf + 1
		} else {
			numsMap[nums[i]] = 1
			single = append(single, nums[i])
		}
	}
	sort.Ints(single)
	for i := 0; i < len(single); i++ {
		middle := single[i]
		goal := -middle
		tempNum := numsMap[middle]
		numsMap[middle] = tempNum - 1
		for j := i; j >= 0; j-- {
			tempL := numsMap[single[j]]
			if tempL > 0 {
				numsMap[single[j]] = tempL - 1
				tempR, okR := numsMap[goal-single[j]]
				if okR && tempR > 0 {
					if single[j] <= middle && middle <= goal-single[j] {
						tempTriple := [3]int{single[j], middle, goal - single[j]}
						res = append(res, tempTriple[:])
					}
				}
				numsMap[single[j]] = tempL
			}
		}
		numsMap[middle] = tempNum
	}
	return res
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}
