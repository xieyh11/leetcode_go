package main

import (
	"../common"
)

type ListNode = common.ListNode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tempAt := head.Next
	head.Next, tempAt.Next = tempAt.Next, head
	head, tempAt = tempAt, head
	for tempAt.Next != nil && tempAt.Next.Next != nil {
		tempAtNext := tempAt.Next
		tempAtNextNext := tempAtNext.Next
		tempAtNext.Next, tempAtNextNext.Next, tempAt.Next = tempAtNextNext.Next, tempAtNext, tempAtNextNext
		tempAt = tempAtNext
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	head.PrintList()

	head = swapPairs(head)
	head.PrintList()
}
