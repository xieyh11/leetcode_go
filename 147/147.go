package main

import (
	"../common"
)

type ListNode = common.ListNode

func insert(head *ListNode, ele *ListNode) *ListNode {
	if head == nil {
		ele.Next = nil
		return ele
	}
	nowAt := head
	former := nowAt
	for nowAt != nil && nowAt.Val < ele.Val {
		nowAt, former = nowAt.Next, nowAt
	}
	if nowAt == head {
		ele.Next = head
		head = ele
	} else if nowAt == nil {
		former.Next, ele.Next = ele, nil
	} else {
		former.Next, ele.Next = ele, former.Next
	}
	return head
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	nowAt := head
	newHead := (*ListNode)(nil)
	for nowAt != nil {
		temp := nowAt.Next
		newHead = insert(newHead, nowAt)
		nowAt = temp
	}
	return newHead
}

func main() {
	head := &ListNode{4, &ListNode{2, &ListNode{1, &ListNode{3, nil}}}}
	head.PrintList()
	head = insertionSortList(head)
	head.PrintList()
}
