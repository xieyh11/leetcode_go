package main

import (
	"../common"
	"fmt"
	"math/rand"
)

type ListNode = common.ListNode

var Intervals = 1000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	Intervals []*ListNode
	Len       int
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	res := Solution{}
	if head == nil {
		return res
	}
	interval := 1
	for head != nil {
		if interval == 1 {
			res.Intervals = append(res.Intervals, head)
			interval = Intervals
		} else {
			interval--
		}
		res.Len++
		head = head.Next
	}
	return res
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	if this.Len == 0 {
		return 0
	}
	idx := rand.Intn(this.Len)
	head := this.Intervals[idx/Intervals]
	remain := idx % Intervals
	for remain != 0 {
		head = head.Next
		remain--
	}
	return head.Val
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}
	sol := Constructor(head)
	for i := 0; i < 10; i++ {
		fmt.Println(sol.GetRandom())
	}
}
