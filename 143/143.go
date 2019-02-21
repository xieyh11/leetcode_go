package main

import (
	"../common"
)

type ListNode = common.ListNode

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	nowAt := head
	stack := []*ListNode{}
	for nowAt != nil {
		stack = append(stack, nowAt)
		nowAt = nowAt.Next
	}
	for i, j := 0, len(stack)-1; i <= j; i, j = i+1, j-1 {
		stack[i].Next, stack[j].Next = stack[j], stack[i].Next
		if i == j {
			stack[i].Next = nil
		}
		if i == j-1 {
			stack[j].Next = nil
		}
	}
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head.PrintList()
	reorderList(head)
	head.PrintList()
}
