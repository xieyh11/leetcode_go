package main

import (
	"../common"
)

type ListNode = common.ListNode

func removeElements(head *ListNode, val int) *ListNode {
	for head != nil && head.Val == val {
		head = head.Next
	}
	if head == nil {
		return head
	}
	former := head
	nowat := head.Next
	for nowat != nil {
		if nowat.Val == val {
			former.Next = nowat.Next
		} else {
			former = nowat
		}
		nowat = nowat.Next
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}
	head.PrintList()
	head = removeElements(head, 1)
	head.PrintList()
}
