package main

import (
	"../common"
)

type ListNode = common.ListNode

type MinHeap struct {
	Heap []*ListNode
	Len  int
}

func (a *MinHeap) Insert(src *ListNode) {
	if a.Len+1 < len(a.Heap) {
		a.Heap[a.Len] = src
	} else {
		a.Heap = append(a.Heap, src)
	}
	a.Len++
	a.MoveUp(a.Len - 1)
}

func (a *MinHeap) MoveUp(b int) {
	i := b
	for i > 0 {
		j := (i - 1) / 2
		if a.Heap[j].Val > a.Heap[i].Val {
			a.Heap[i], a.Heap[j] = a.Heap[j], a.Heap[i]
			i = j
		} else {
			break
		}
	}
}

func (a *MinHeap) MoveDown(b int) {
	i := b
	for i < a.Len {
		il := i*2 + 1
		ir := i*2 + 2
		if ir < a.Len {
			minI := i
			if a.Heap[minI].Val > a.Heap[il].Val {
				minI = il
			}
			if a.Heap[minI].Val > a.Heap[ir].Val {
				minI = ir
			}
			if i != minI {
				a.Heap[i], a.Heap[minI] = a.Heap[minI], a.Heap[i]
				i = minI
			} else {
				break
			}
		} else {
			if il < a.Len {
				if a.Heap[i].Val > a.Heap[il].Val {
					a.Heap[i], a.Heap[il] = a.Heap[il], a.Heap[i]
				}
			}
			break
		}
	}
}

func (a *MinHeap) ExtractMin() *ListNode {
	a.Heap[a.Len-1], a.Heap[0] = a.Heap[0], a.Heap[a.Len-1]
	a.Len--
	a.MoveDown(0)
	return a.Heap[a.Len]
}

func (a *MinHeap) Enlarge(idx int, val *ListNode) {
	a.Heap[idx] = val
	a.MoveDown(idx)
}

func mergeKLists(lists []*ListNode) *ListNode {
	myHeap := MinHeap{}
	for _, head := range lists {
		if head != nil {
			myHeap.Insert(head)
		}
	}
	head := (*ListNode)(nil)
	tempAt := head
	for myHeap.Len > 1 {
		if head == nil {
			head = myHeap.Heap[0]
			tempAt = head
		} else {
			tempAt.Next = myHeap.Heap[0]
			tempAt = tempAt.Next
		}
		temp := myHeap.Heap[0].Next
		if temp == nil {
			myHeap.ExtractMin()
		} else {
			myHeap.Enlarge(0, temp)
		}
	}

	if myHeap.Len == 1 {
		if head == nil {
			head = myHeap.Heap[0]
		} else {
			tempAt.Next = myHeap.Heap[0]
		}
	}

	return head
}

func main() {
	l0 := &ListNode{1, &ListNode{4, &ListNode{5, nil}}}
	l1 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}
	l2 := &ListNode{2, &ListNode{6, nil}}
	l0.PrintList()
	lists := []*ListNode{l0, l1, l2}
	head := mergeKLists(lists)
	head.PrintList()
}
