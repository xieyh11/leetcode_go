package main

import (
	"fmt"
)

type PairSum struct {
	I   [2]int
	Sum int
}

type Heap struct {
	Pairs []PairSum
}

func (this *Heap) Add(pair [2]int, sum int) {
	this.Pairs = append(this.Pairs, PairSum{pair, sum})
	this.moveUp(len(this.Pairs) - 1)
}

func (this *Heap) RemoveTop() ([2]int, int) {
	res := this.Pairs[0]
	last := len(this.Pairs) - 1
	this.Pairs[0] = this.Pairs[last]
	this.Pairs = this.Pairs[:last]
	this.moveDown(0)
	return res.I, res.Sum
}

func (this *Heap) moveUp(idx int) {
	parent := (idx - 1) / 2
	for parent >= 0 && this.Pairs[parent].Sum > this.Pairs[idx].Sum {
		this.Pairs[parent], this.Pairs[idx] = this.Pairs[idx], this.Pairs[parent]
		parent, idx = (parent-1)/2, parent
	}
}

func (this *Heap) moveDown(idx int) {
	left, right := idx*2+1, idx*2+2
	heapLen := len(this.Pairs)
	for left < heapLen {
		min := left
		if right < heapLen && this.Pairs[right].Sum < this.Pairs[min].Sum {
			min = right
		}
		if this.Pairs[idx].Sum > this.Pairs[min].Sum {
			this.Pairs[idx], this.Pairs[min] = this.Pairs[min], this.Pairs[idx]
			left, right, idx = min*2+1, min*2+2, min
		} else {
			break
		}
	}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return [][]int{}
	}
	nums1Len, nums2Len := len(nums1), len(nums2)
	k = getMin(k, nums1Len*nums2Len)
	searched := make(map[[2]int]bool)
	heap := Heap{}
	pair := [2]int{0, 0}
	pairSum := nums1[pair[0]] + nums2[pair[1]]
	heap.Add(pair, pairSum)
	res := make([][]int, 0)
	for i := 0; i < k; i++ {
		minPair, _ := heap.RemoveTop()
		res = append(res, []int{nums1[minPair[0]], nums2[minPair[1]]})
		next1, next2 := [2]int{minPair[0] + 1, minPair[1]}, [2]int{minPair[0], minPair[1] + 1}

		if next1[0] < nums1Len {
			if _, ok := searched[next1]; !ok {
				tempSum := nums1[next1[0]] + nums2[next1[1]]
				heap.Add(next1, tempSum)
				searched[next1] = true
			}
		}

		if next2[1] < nums2Len {
			if _, ok := searched[next2]; !ok {
				tempSum := nums1[next2[0]] + nums2[next2[1]]
				heap.Add(next2, tempSum)
				searched[next2] = true
			}
		}
	}
	return res
}

func main() {
	nums1 := []int{1, 7, 11}
	nums2 := []int{2, 4, 6}
	k := 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 1, 2}
	nums2 = []int{1, 2, 3}
	k = 2
	fmt.Println(kSmallestPairs(nums1, nums2, k))

	nums1 = []int{1, 2}
	nums2 = []int{3}
	k = 3
	fmt.Println(kSmallestPairs(nums1, nums2, k))
}
