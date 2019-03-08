package main

import (
	"../common"
)

type ListNode = common.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	former, nowat := head, head.Next
	head.Next = nil
	for nowat != nil {
		temp := nowat.Next
		nowat.Next = former
		former, nowat = nowat, temp
	}
	return former
}

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	head.PrintList()
	head = reverseList(head)
	head.PrintList()
}
