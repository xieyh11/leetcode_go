package main

import (
	"fmt"
)

func recurCombine(array []int, k int) [][]int {
	n := len(array)
	if n < k || k == 0 {
		return [][]int{}
	}
	if n == k {
		return [][]int{array}
	}
	if k == 1 {
		res := make([][]int, n)
		for i := range res {
			res[i] = []int{array[i]}
		}
		return res
	}
	mid := n / 2
	res := make([][]int, 0)
	for tempK := 0; tempK <= k && tempK <= mid; tempK++ {
		if k-tempK <= n-mid {
			left := recurCombine(array[:mid], tempK)
			right := recurCombine(array[mid:], k-tempK)
			if tempK == 0 {
				res = append(res, right...)
			} else if k == tempK {
				res = append(res, left...)
			} else {
				for _, l := range left {
					for _, r := range right {
						tempR := make([]int, k)
						copy(tempR[:tempK], l)
						copy(tempR[tempK:], r)
						res = append(res, tempR)
					}
				}
			}
		}
	}
	return res
}

func combine(n int, k int) [][]int {
	array := make([]int, n)
	for i := range array {
		array[i] = i + 1
	}
	return recurCombine(array, k)
}

func main() {
	fmt.Println(combine(4, 2))
}
