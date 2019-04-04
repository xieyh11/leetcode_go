package main

import (
	"fmt"
)

type NumArray struct {
	Nums []int
	Sums []int
}

func Constructor(nums []int) NumArray {
	sums := make([]int, len(nums)+1)
	for i := range nums {
		sums[i+1] = sums[i] + nums[i]
	}
	return NumArray{nums, sums}
}

func (this *NumArray) Update(i int, val int) {
	if this.Nums[i] == val {
		return
	}
	diff := val - this.Nums[i]
	this.Nums[i] = val
	for j := i + 1; j < len(this.Sums); j++ {
		this.Sums[j] += diff
	}
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.Sums[j+1] - this.Sums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */

func main() {
	nums := []int{1, 3, 5}
	sum := Constructor(nums)
	fmt.Println(sum.SumRange(0, 2))
	sum.Update(1, 2)
	fmt.Println(sum.SumRange(0, 2))
}
