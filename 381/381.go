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
func ConstructorSet() *RandomizedSet {
	return &RandomizedSet{make([]int, 0), make(map[int]int)}
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

func (this *RandomizedSet) Len() int {
	return len(this.Nums)
}

type RandomizedCollection struct {
	Nums   []int
	Indexs map[int]*RandomizedSet
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{make([]int, 0), make(map[int]*RandomizedSet)}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.Nums = append(this.Nums, val)
	_, res := this.Indexs[val]
	if !res {
		this.Indexs[val] = ConstructorSet()
		res = !res
	} else {
		res = this.Indexs[val].Len() == 0
	}
	this.Indexs[val].Insert(len(this.Nums) - 1)
	return res
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	v, res := this.Indexs[val]
	if !res || v.Len() == 0 {
		return false
	}
	vIdx := v.GetRandom()
	v.Remove(vIdx)
	last := len(this.Nums) - 1
	if vIdx != last {
		lastV := this.Nums[last]
		this.Nums[vIdx] = this.Nums[last]
		this.Indexs[lastV].Remove(last)
		this.Indexs[lastV].Insert(vIdx)
	}
	this.Nums = this.Nums[:last]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.Nums[rand.Intn(len(this.Nums))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */

func main() {
	collection := Constructor()
	fmt.Println(collection.Insert(0))
	fmt.Println(collection.Remove(0))
	fmt.Println(collection.Insert(-1))
	fmt.Println(collection.Remove(0))
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
	fmt.Println(collection.GetRandom())
}
