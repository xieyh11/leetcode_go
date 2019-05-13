package main

import (
	"fmt"
	"math/rand"
)

type RandomizedSet struct {
	Nums     []int
	IndexMap map[int]int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 0), make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.IndexMap[val]; ok {
		return false
	}
	this.Nums = append(this.Nums, val)
	this.IndexMap[val] = len(this.Nums) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if v, ok := this.IndexMap[val]; ok {
		delete(this.IndexMap, val)
		last := len(this.Nums) - 1
		lastV := this.Nums[last]
		if v != last {
			this.Nums[v] = lastV
			this.IndexMap[lastV] = v
		}
		this.Nums = this.Nums[:last]
		return true
	}
	return false
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.Nums[rand.Intn(len(this.Nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	randomSet := Constructor()
	fmt.Println(randomSet.Insert(1))
	fmt.Println(randomSet.Remove(2))
	fmt.Println(randomSet.GetRandom())
	fmt.Println(randomSet.Insert(2))
}
