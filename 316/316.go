package main

import (
	"container/list"
	"fmt"
)

type HeapNode struct {
	Val   byte
	Index int
}

type MinHeap struct {
	Heap     []HeapNode
	IndexMap map[HeapNode]int
}

func Constructor() MinHeap {
	return MinHeap{IndexMap: make(map[HeapNode]int)}
}

func moveDown(heap []HeapNode, heapMap map[HeapNode]int, idx int) {
	childLeft := idx*2 + 1
	childRight := childLeft + 1
	heapLen := len(heap)
	for (childLeft < heapLen && heap[idx].Index > heap[childLeft].Index) || (childRight < heapLen && heap[idx].Index > heap[childRight].Index) {
		minChild := childLeft
		if childRight < heapLen && heap[minChild].Index > heap[childRight].Index {
			minChild = childRight
		}
		heapMap[heap[idx]], heapMap[heap[minChild]] = minChild, idx
		heap[idx], heap[minChild] = heap[minChild], heap[idx]
		idx, childLeft = minChild, minChild*2+1
		childRight = childLeft + 1
	}
}

func moveUp(heap []HeapNode, heapMap map[HeapNode]int, idx int) {
	parent := (idx - 1) / 2
	for parent >= 0 && heap[parent].Index > heap[idx].Index {
		heapMap[heap[parent]], heapMap[heap[idx]] = idx, parent
		heap[parent], heap[idx] = heap[idx], heap[parent]
		parent, idx = (parent-1)/2, parent
	}
}

func (this *MinHeap) Insert(node HeapNode) {
	this.Heap = append(this.Heap, node)
	this.IndexMap[node] = len(this.Heap) - 1
	moveUp(this.Heap, this.IndexMap, len(this.Heap)-1)
}

func (this *MinHeap) GetMin() HeapNode {
	return this.Heap[0]
}

func (this *MinHeap) RemoveMin() HeapNode {
	res := this.Heap[0]
	this.Heap[0] = this.Heap[len(this.Heap)-1]
	this.Heap = this.Heap[:len(this.Heap)-1]
	this.IndexMap[this.Heap[0]] = 0
	delete(this.IndexMap, res)
	if len(this.Heap) > 0 {
		moveDown(this.Heap, this.IndexMap, 0)
	}
	return res
}

func (this *MinHeap) RemoveNode(node HeapNode) {
	loc := this.IndexMap[node]
	if loc == len(this.Heap)-1 {
		delete(this.IndexMap, node)
		this.Heap = this.Heap[:len(this.Heap)-1]
	} else {
		this.Heap[loc] = this.Heap[len(this.Heap)-1]
		this.IndexMap[this.Heap[loc]] = loc
		delete(this.IndexMap, node)
		this.Heap = this.Heap[:len(this.Heap)-1]
		if len(this.Heap) > 0 {
			moveDown(this.Heap, this.IndexMap, loc)
			moveUp(this.Heap, this.IndexMap, loc)
		}
	}
}

func recSliceSol(vals []int, left, right int, num int) []int {
	if left == right {
		if vals[left] > num {
			return vals[left:]
		}
		return vals[left+1:]
	}
	if left == right+1 {
		if vals[left] > num {
			return vals[left:]
		}
		if vals[right] > num {
			return vals[right:]
		}
		return vals[right+1:]
	}

	mid := (left + right) / 2
	if vals[mid] == num {
		return vals[mid+1:]
	} else if vals[mid] > num {
		return recSliceSol(vals, left, mid-1, num)
	} else {
		return recSliceSol(vals, mid+1, right, num)
	}
}

func sliceLarger(vals []int, num int) []int {
	return recSliceSol(vals, 0, len(vals)-1, num)
}

func removeDuplicateLetters(s string) string {
	letterLoc := make(map[byte][]int, 26)
	for i := range s {
		letterLoc[s[i]] = append(letterLoc[s[i]], i)
	}
	res := make([]byte, 0)
	lastLoc := Constructor()
	letters := list.New()
	for k, v := range letterLoc {
		lastLoc.Insert(HeapNode{k, v[len(v)-1]})
	}
	for l := byte('a'); l <= 'z'; l++ {
		if _, ok := letterLoc[l]; ok {
			letters.PushBack(l)
		}
	}

	for letters.Len() > 0 {
		now := letters.Front()
		minLast := lastLoc.GetMin()
		for now != nil {
			l := now.Value.(byte)
			lLoc := letterLoc[l]
			if minLast.Val == l || minLast.Index > lLoc[0] {
				res = append(res, l)
				lastLoc.RemoveNode(HeapNode{l, lLoc[len(lLoc)-1]})
				delete(letterLoc, l)
				for k, v := range letterLoc {
					letterLoc[k] = sliceLarger(v, lLoc[0])
				}
				break
			} else {
				now = now.Next()
			}
		}
		letters.Remove(now)
	}
	return string(res)
}

func main() {
	s := "aaa"
	fmt.Println(removeDuplicateLetters(s))
}
