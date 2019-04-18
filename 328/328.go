package main

import (
	"../common"
)

type ListNode = common.ListNode

func oddEvenList(head *ListNode) *ListNode {
	odd := (*ListNode)(nil)
	even := odd
	evenHead := even
	num := 1
	nowAt := head
	for nowAt != nil {
		if num&1 == 0 {
			if even == nil {
				evenHead = nowAt
			} else {
				even.Next = nowAt
			}
			even = nowAt
		} else {
			if odd != nil {
				odd.Next = nowAt
			}
			odd = nowAt
		}
		num++
		nowAt = nowAt.Next
	}
	if odd != nil {
		odd.Next = evenHead
	}
	if even != nil {
		even.Next = nil
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head = oddEvenList(head)
	head.PrintList()
}
