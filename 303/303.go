package main

import (
	"fmt"
)

type NumArray struct {
	Sums []int
}

func Constructor(nums []int) NumArray {
	res := NumArray{make([]int, len(nums))}
	if len(nums) > 0 {
		res.Sums[0] = nums[0]
		for i := 1; i < len(nums); i++ {
			res.Sums[i] = res.Sums[i-1] + nums[i]
		}
	}
	return res
}

func (this *NumArray) SumRange(i int, j int) int {
	if i < 0 {
		i = 0
	}
	if j >= len(this.Sums) {
		j = len(this.Sums) - 1
	}
	res := this.Sums[j]
	if i > 0 {
		res -= this.Sums[i-1]
	}
	return res
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	nums := []int{-2, 0, 3, -5, 2, -1}
	array := Constructor(nums)
	fmt.Println(array.SumRange(0, 2))
	fmt.Println(array.SumRange(2, 5))
	fmt.Println(array.SumRange(0, 5))
}
