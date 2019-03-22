package main

import (
	"container/list"
	"fmt"
)

type MyQueue struct {
	S1 *list.List
	S2 *list.List
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{list.New(), list.New()}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.S1.PushBack(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if this.S2.Len() == 0 {
		for this.S1.Len() > 0 {
			this.S2.PushBack(this.S1.Remove(this.S1.Back()))
		}
	}
	return this.S2.Remove(this.S2.Back()).(int)
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if this.S2.Len() == 0 {
		for this.S1.Len() > 0 {
			this.S2.PushBack(this.S1.Remove(this.S1.Back()))
		}
	}
	return this.S2.Back().Value.(int)
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return this.S1.Len() == 0 && this.S2.Len() == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	queue := Constructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	queue.Push(3)
	fmt.Println(queue.Pop())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}
