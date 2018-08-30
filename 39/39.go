package main

import (
	"fmt"
	"sort"
)

func repInt(num, rep int) []int {
	res := make([]int, rep)
	for i := range res {
		res[i] = num
	}
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	if len(candidates) == 1 {
		if target%candidates[0] == 0 {
			return [][]int{repInt(candidates[0], target/candidates[0])}
		} else {
			return [][]int{}
		}
	}
	res := make([][]int, 0)
	for i := 0; i <= target/candidates[0]; i++ {
		tempRes := combinationSum(candidates[1:], target-candidates[0]*i)
		if len(tempRes) > 0 {
			tempRep := repInt(candidates[0], i)
			for _, v := range tempRes {
				res = append(res, append(tempRep, v...))
			}
		}
	}
	return res
}

func main() {
	nums := []int{2, 3, 5}
	target := 8
	fmt.Println(combinationSum(nums, target))
}
