package main

import (
	"../common"
)

type ListNode = common.ListNode

func partition(head *ListNode, x int) *ListNode {
	left := (*ListNode)(nil)
	right := (*ListNode)(nil)
	headL := left
	headR := right
	for head != nil {
		if head.Val < x {
			if headL == nil {
				headL = head
				left = headL
			} else {
				left.Next = head
				left = left.Next
			}
		} else {
			if headR == nil {
				headR = head
				right = headR
			} else {
				right.Next = head
				right = right.Next
			}
		}
		head = head.Next
	}
	if right != nil {
		right.Next = nil
	}
	if headL == nil {
		return headR
	} else {
		left.Next = headR
		return headL
	}
}

func main() {
	head := &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}}
	head.PrintList()
	head = partition(head, 3)
	head.PrintList()
}
