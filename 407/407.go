package main

import (
	"fmt"
)

type HeapNode struct {
	Loc [2]int
	Val int
}

type LocHeap struct {
	Heap  []HeapNode
	Index map[[2]int]int
}

func (this *LocHeap) Add(loc [2]int, val int) {
	this.Heap = append(this.Heap, HeapNode{loc, val})
	idx := len(this.Heap) - 1
	this.Index[loc] = idx
	this.moveUp(idx)
}

func (this *LocHeap) RemoveTop() ([2]int, int) {
	res := this.Heap[0]
	last := len(this.Heap) - 1
	this.Heap[0] = this.Heap[last]
	this.Index[this.Heap[0].Loc] = 0
	this.Heap = this.Heap[:last]
	delete(this.Index, res.Loc)
	this.moveDown(0)
	return res.Loc, res.Val
}

func (this *LocHeap) InsertOrUpdateToMin(loc [2]int, val int) {
	idx, ok := this.Index[loc]
	if !ok {
		this.Add(loc, val)
	} else if val < this.Heap[idx].Val {
		this.Heap[idx].Val = val
		this.moveUp(idx)
	}
}

func (this *LocHeap) Empty() bool {
	return len(this.Heap) == 0
}

func (this *LocHeap) moveUp(idx int) {
	parent := (idx - 1) / 2
	for parent >= 0 && this.Heap[parent].Val > this.Heap[idx].Val {
		this.Heap[parent], this.Heap[idx] = this.Heap[idx], this.Heap[parent]
		this.Index[this.Heap[idx].Loc], this.Index[this.Heap[parent].Loc] = idx, parent
		parent, idx = (parent-1)/2, parent
	}
}

func (this *LocHeap) moveDown(idx int) {
	left, right := idx*2+1, idx*2+2
	heapLen := len(this.Heap)
	for left < heapLen {
		min := left
		if right < heapLen && this.Heap[right].Val < this.Heap[min].Val {
			min = right
		}
		if this.Heap[idx].Val > this.Heap[min].Val {
			this.Heap[idx], this.Heap[min] = this.Heap[min], this.Heap[idx]
			this.Index[this.Heap[idx].Loc], this.Index[this.Heap[min].Loc] = idx, min
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

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}
	rows, cols := len(heightMap), len(heightMap[0])
	if rows <= 2 || cols <= 2 {
		return 0
	}
	res := 0
	heap := LocHeap{make([]HeapNode, 0), make(map[[2]int]int)}

	searched := make([][]bool, rows)
	for i := range searched {
		searched[i] = make([]bool, cols)
	}

	for i := 1; i < cols-1; i++ {
		heap.InsertOrUpdateToMin([2]int{1, i}, heightMap[0][i])
		heap.InsertOrUpdateToMin([2]int{rows - 2, i}, heightMap[rows-1][i])
	}

	for j := 1; j < rows-1; j++ {
		heap.InsertOrUpdateToMin([2]int{j, 1}, heightMap[j][0])
		heap.InsertOrUpdateToMin([2]int{j, cols - 2}, heightMap[j][cols-1])
	}

	for !heap.Empty() {
		minIdx, minVal := heap.RemoveTop()
		searched[minIdx[0]][minIdx[1]] = true
		if minVal > heightMap[minIdx[0]][minIdx[1]] {
			res += minVal - heightMap[minIdx[0]][minIdx[1]]
		} else {
			minVal = heightMap[minIdx[0]][minIdx[1]]
		}
		if minIdx[0]-1 > 0 && !searched[minIdx[0]-1][minIdx[1]] {
			heap.InsertOrUpdateToMin([2]int{minIdx[0] - 1, minIdx[1]}, minVal)
		}
		if minIdx[0]+1 < rows-1 && !searched[minIdx[0]+1][minIdx[1]] {
			heap.InsertOrUpdateToMin([2]int{minIdx[0] + 1, minIdx[1]}, minVal)
		}

		if minIdx[1]-1 > 0 && !searched[minIdx[0]][minIdx[1]-1] {
			heap.InsertOrUpdateToMin([2]int{minIdx[0], minIdx[1] - 1}, minVal)
		}
		if minIdx[1]+1 < cols-1 && !searched[minIdx[0]][minIdx[1]+1] {
			heap.InsertOrUpdateToMin([2]int{minIdx[0], minIdx[1] + 1}, minVal)
		}
	}
	return res
}

func main() {
	height := [][]int{{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1}}
	fmt.Println(trapRainWater(height))
}
