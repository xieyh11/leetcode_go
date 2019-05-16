package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	Index map[int][]int
}

func Constructor(nums []int) Solution {
	index := make(map[int][]int)
	for i, v := range nums {
		index[v] = append(index[v], i)
	}
	return Solution{index}
}

func (this *Solution) Pick(target int) int {
	if v, ok := this.Index[target]; ok {
		return v[rand.Intn(len(v))]
	}
	return -1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */

func main() {
	nums := []int{1, 2, 3, 3, 3}
	sol := Constructor(nums)
	fmt.Println(sol.Pick(3))
	fmt.Println(sol.Pick(2))
}
