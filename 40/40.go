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

func recursiveSol(candidates []int, target int, numOf map[int]int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	if len(candidates) <= numOf[candidates[0]] {
		if target%candidates[0] == 0 && target/candidates[0] <= len(candidates) {
			return [][]int{repInt(candidates[0], target/candidates[0])}
		} else {
			return [][]int{}
		}
	}
	res := make([][]int, 0)
	for i := 0; i <= target/candidates[0] && i <= numOf[candidates[0]]; i++ {
		tempRes := recursiveSol(candidates[numOf[candidates[0]]:], target-candidates[0]*i, numOf)
		if len(tempRes) > 0 {
			tempRep := repInt(candidates[0], i)
			for _, v := range tempRes {
				res = append(res, append(tempRep, v...))
			}
		}
	}
	return res
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	numOf := make(map[int]int)
	for _, v := range candidates {
		if _, ok := numOf[v]; ok {
			numOf[v]++
		} else {
			numOf[v] = 1
		}
	}
	return recursiveSol(candidates, target, numOf)
}

func main() {
	nums := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(nums, target))
}
