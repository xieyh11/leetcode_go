package main

import (
	"fmt"
)

type Pos struct {
	Row int
	Val int
}

type Heap struct {
	Heap []Pos
}

func (this *Heap) Add(row, val int) {
	this.Heap = append(this.Heap, Pos{row, val})
	this.moveUp(len(this.Heap) - 1)
}

func (this *Heap) RemoveTop() (int, int) {
	res := this.Heap[0]
	last := len(this.Heap) - 1
	this.Heap[0] = this.Heap[last]
	this.Heap = this.Heap[:last]
	this.moveDown(0)
	return res.Row, res.Val
}

func (this *Heap) moveUp(idx int) {
	parent := (idx - 1) / 2
	for parent >= 0 && this.Heap[parent].Val > this.Heap[idx].Val {
		this.Heap[parent], this.Heap[idx] = this.Heap[idx], this.Heap[parent]
		parent, idx = (parent-1)/2, parent
	}
}

func (this *Heap) moveDown(idx int) {
	left, right := idx*2+1, idx*2+2
	heapLen := len(this.Heap)
	for left < heapLen {
		min := left
		if right < heapLen && this.Heap[right].Val < this.Heap[min].Val {
			min = right
		}
		if this.Heap[idx].Val > this.Heap[min].Val {
			this.Heap[idx], this.Heap[min] = this.Heap[min], this.Heap[idx]
			left, right, idx = min*2+1, min*2+2, min
		} else {
			break
		}
	}
}

func kthSmallest(matrix [][]int, k int) int {
	if len(matrix) == 0 {
		return 0
	}
	n := len(matrix)
	rowsIndex := make([]int, n)
	heap := Heap{}
	heap.Add(0, matrix[0][0])
	for k > 1 {
		row, _ := heap.RemoveTop()
		k--
		rowsIndex[row]++
		if rowsIndex[row] < n && (row == 0 || rowsIndex[row] < rowsIndex[row-1]) {
			heap.Add(row, matrix[row][rowsIndex[row]])
		}
		if row < n-1 && rowsIndex[row+1] == rowsIndex[row]-1 {
			heap.Add(row+1, matrix[row+1][rowsIndex[row+1]])
		}
	}
	_, val := heap.RemoveTop()
	return val
}

func main() {
	matrix := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	for k := 1; k <= 9; k++ {
		fmt.Println(kthSmallest(matrix, k))
	}
}
