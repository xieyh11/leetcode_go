package main

import (
	"fmt"
	"math/rand"
)

type NumFreq struct {
	Num  int
	Freq int
}

func (a *NumFreq) Less(b NumFreq) bool {
	return a.Freq < b.Freq
}

func (a *NumFreq) Equal(b NumFreq) bool {
	return a.Freq == b.Freq
}

func (a *NumFreq) LessEq(b NumFreq) bool {
	return a.Freq <= b.Freq
}

func findKth(nums []NumFreq, k int) []int {
	if len(nums) == 1 {
		return []int{nums[0].Num}
	}
	if len(nums) == 2 {
		smaller, larger := nums[0], nums[1]
		if larger.Less(smaller) {
			smaller, larger = larger, smaller
		}
		if k == 1 {
			return []int{larger.Num}
		}
		return []int{larger.Num, smaller.Num}
	}
	idx := rand.Intn(len(nums))
	mid := nums[idx]
	former := 0
	latter := len(nums) - 1
	for former < latter {
		if mid.LessEq(nums[former]) {
			former++
		} else if nums[latter].Less(mid) {
			latter--
		} else {
			nums[former], nums[latter] = nums[latter], nums[former]
			former++
			latter--
		}
	}
	if mid.LessEq(nums[former]) {
		former++
	}
	lessOrEqual := former
	former, latter = 0, lessOrEqual-1
	for former < latter {
		if mid.Less(nums[former]) {
			former++
		} else if mid.Equal(nums[latter]) {
			latter--
		} else {
			nums[former], nums[latter] = nums[latter], nums[former]
			former++
			latter--
		}
	}
	if mid.Less(nums[former]) {
		former++
	}
	less := former
	if k <= less {
		return findKth(nums[:less], k)
	}
	fomerInts := make([]int, less)
	for i := 0; i < less; i++ {
		fomerInts[i] = nums[i].Num
	}
	if k <= lessOrEqual {
		for i := less; i < k; i++ {
			fomerInts = append(fomerInts, nums[i].Num)
		}
		return fomerInts
	} else {
		for i := less; i < lessOrEqual; i++ {
			fomerInts = append(fomerInts, nums[i].Num)
		}
		return append(fomerInts, findKth(nums[lessOrEqual:], k-lessOrEqual)...)
	}
}

func topKFrequent(nums []int, k int) []int {
	store := make(map[int]int)
	for _, num := range nums {
		store[num]++
	}
	numFreq := make([]NumFreq, 0)
	for k, v := range store {
		numFreq = append(numFreq, NumFreq{k, v})
	}
	return findKth(numFreq, k)
}

func main() {
	nums0 := []int{1, 1, 1, 2, 2, 3}
	k0 := 2
	fmt.Println(topKFrequent(nums0, k0))
}
