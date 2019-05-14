package main

import (
	"fmt"
	"math/rand"
)

type Solution struct {
	Nums []int
}

func Constructor(nums []int) Solution {
	return Solution{nums}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.Nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.Nums)
	res := make([]int, n)
	copy(res, this.Nums)
	for n > 1 {
		idx := rand.Intn(n)
		res[n-1], res[idx] = res[idx], res[n-1]
		n--
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */

func main() {
	nums := []int{1, 2, 3}
	sol := Constructor(nums)
	for i := 0; i < 10; i++ {
		fmt.Println(sol.Shuffle())
	}
}
