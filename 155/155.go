package main

import (
	"container/list"
	"fmt"
)

type StackNode struct {
	Val int
	Idx int
}

type MinStack struct {
	MinStack []*list.Element
	Stack    *list.List
	Len      int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{make([]*list.Element, 0), list.New(), 0}
}

func minStackUp(stack []*list.Element, start int) {
	nowIdx := start
	for nowIdx > 0 {
		parentIdx := (nowIdx - 1) / 2
		nowNode := stack[nowIdx].Value.(*StackNode)
		parentNode := stack[parentIdx].Value.(*StackNode)
		if parentNode.Val > nowNode.Val {
			stack[parentIdx], stack[nowIdx] = stack[nowIdx], stack[parentIdx]
			parentNode.Idx, nowNode.Idx = nowNode.Idx, parentNode.Idx
			nowIdx = parentIdx
		} else {
			break
		}
	}
}

func minStackDown(stack []*list.Element, stackLen int, start int) {
	nowIdx := start
	for nowIdx < stackLen {
		leftIdx := 2*nowIdx + 1
		rightIdx := 2*nowIdx + 2
		hasLeft := leftIdx < stackLen
		hasRight := rightIdx < stackLen
		if hasLeft || hasRight {
			nowNode := stack[nowIdx].Value.(*StackNode)
			leftNode := (*StackNode)(nil)
			rightNode := (*StackNode)(nil)
			minIdx := -1
			minNode := leftNode
			if hasLeft {
				leftNode = stack[leftIdx].Value.(*StackNode)
				minIdx = leftIdx
				minNode = leftNode
			}
			if hasRight {
				rightNode = stack[rightIdx].Value.(*StackNode)
				if minIdx == -1 || minNode.Val > rightNode.Val {
					minIdx = rightIdx
					minNode = rightNode
				}
			}
			if minNode.Val < nowNode.Val {
				stack[nowIdx], stack[minIdx] = stack[minIdx], stack[nowIdx]
				minNode.Idx, nowNode.Idx = nowNode.Idx, minNode.Idx
				nowIdx = minIdx
			} else {
				break
			}
		} else {
			break
		}
	}
}

func (this *MinStack) Push(x int) {
	newNode := new(StackNode)
	newNode.Val = x
	e := this.Stack.PushBack(newNode)
	if this.Len < len(this.MinStack) {
		this.MinStack[this.Len] = e
		this.Len++
	} else {
		this.Len++
		this.MinStack = append(this.MinStack, e)
	}
	newNode.Idx = this.Len - 1
	minStackUp(this.MinStack, newNode.Idx)
}

func (this *MinStack) Pop() {
	this.Len--
	e := this.Stack.Back()
	lastIdx := this.Len
	lastNode := this.MinStack[lastIdx].Value.(*StackNode)
	eNode := e.Value.(*StackNode)
	this.MinStack[eNode.Idx] = this.MinStack[lastIdx]
	lastNode.Idx = eNode.Idx
	minStackDown(this.MinStack, this.Len, lastNode.Idx)
	this.Stack.Remove(e)
}

func (this *MinStack) Top() int {
	return this.Stack.Back().Value.(*StackNode).Val
}

func (this *MinStack) GetMin() int {
	return this.MinStack[0].Value.(*StackNode).Val
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	stack := Constructor()
	stack.Push(-10)
	stack.Push(14)
	fmt.Println(stack.GetMin())
	fmt.Println(stack.GetMin())
	stack.Push(-20)
	fmt.Println(stack.GetMin())
	fmt.Println(stack.GetMin())
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Pop()
	stack.Push(10)
	stack.Push(-7)
	fmt.Println(stack.GetMin())
	stack.Push(-7)
	stack.Pop()
	fmt.Println(stack.Top())
	fmt.Println(stack.GetMin())
	stack.Pop()
}
