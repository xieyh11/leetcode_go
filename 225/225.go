package main

import (
	"container/list"
	"fmt"
)

type MyStack struct {
	Queue *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	size := this.Queue.Len()
	this.Queue.PushBack(x)
	for size > 0 {
		size--
		this.Queue.MoveToBack(this.Queue.Front())
	}
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.Queue.Remove(this.Queue.Front()).(int)
}

/** Get the top element. */
func (this *MyStack) Top() int {
	return this.Queue.Front().Value.(int)
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.Queue.Len() == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	stack := Constructor()
	stack.Push(1)
	stack.Push(2)
	fmt.Println(stack.Top())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Top())
	fmt.Println(stack.Empty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Empty())
}
