package main

import (
	"fmt"
	"sort"
)

func duplicK(num, times int) []int {
	res := make([]int, times)
	for i := range res {
		res[i] = num
	}
	return res
}

func nChooseK(n, k int) [][]int {
	if k == 0 {
		return [][]int{make([]int, n)}
	}
	if n < k {
		return [][]int{}
	}
	if n == k {
		return [][]int{duplicK(1, n)}
	}
	res := make([][]int, 0)
	if k == 1 {
		for i := 0; i < n; i++ {
			temp := make([]int, n)
			temp[i] = 1
			res = append(res, temp)
		}
		return res
	}

	left := n / 2
	right := n - left
	for ki := 0; ki <= left && ki <= k; ki++ {
		tempL := nChooseK(left, ki)
		tempR := nChooseK(right, k-ki)
		for _, l := range tempL {
			for _, r := range tempR {
				tempRes := make([]int, n)
				copy(tempRes[:left], l)
				copy(tempRes[left:], r)
				res = append(res, tempRes)
			}
		}
	}
	return res
}

func recursionSol(nums []int, numOf map[int]int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	if len(nums) <= numOf[nums[0]] {
		return [][]int{duplicK(nums[0], len(nums))}
	}

	res := make([][]int, 0)
	numF := numOf[nums[0]]
	tempRes := recursionSol(nums[numF:], numOf)
	n := len(nums)
	locs := nChooseK(n, numF)
	for _, loc := range locs {
		for _, per := range tempRes {
			oneRes := make([]int, n)
			copy(oneRes, loc)
			peri := 0
			for i := range oneRes {
				if oneRes[i] == 1 {
					oneRes[i] = nums[0]
				} else {
					oneRes[i] = per[peri]
					peri++
				}
			}
			res = append(res, oneRes)
		}
	}
	return res
}

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)

	numOf := make(map[int]int)
	for _, v := range nums {
		if _, ok := numOf[v]; ok {
			numOf[v]++
		} else {
			numOf[v] = 1
		}
	}

	return recursionSol(nums, numOf)
}

func main() {
	nums := []int{1, 1, 2}
	fmt.Println(permuteUnique(nums))
}
