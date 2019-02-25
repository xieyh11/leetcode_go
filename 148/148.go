package main

import (
	"../common"
)

type ListNode = common.ListNode

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Val > head.Next.Val {
			head, head.Next, head.Next.Next = head.Next, nil, head
		}
		return head
	}
	one, two := head, head
	for two.Next != nil && two.Next.Next != nil {
		one = one.Next
		two = two.Next.Next
	}
	nextHead := one.Next
	one.Next = nil
	left := sortList(head)
	right := sortList(nextHead)
	newHead := (*ListNode)(nil)
	nowAt := newHead
	for left != nil && right != nil {
		if left.Val <= right.Val {
			temp := left.Next
			if newHead == nil {
				newHead = left
				nowAt = newHead
			} else {
				nowAt.Next = left
				nowAt = nowAt.Next
			}
			left = temp
		} else {
			temp := right.Next
			if newHead == nil {
				newHead = right
				nowAt = newHead
			} else {
				nowAt.Next = right
				nowAt = nowAt.Next
			}
			right = temp
		}
	}
	for left != nil {
		temp := left.Next
		if newHead == nil {
			newHead = left
			nowAt = newHead
		} else {
			nowAt.Next = left
			nowAt = nowAt.Next
		}
		left = temp
	}
	for right != nil {
		temp := right.Next
		if newHead == nil {
			newHead = right
			nowAt = newHead
		} else {
			nowAt.Next = right
			nowAt = nowAt.Next
		}
		right = temp
	}
	nowAt.Next = nil
	return newHead
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head.PrintList()
	head = sortList(head)
	head.PrintList()
}
