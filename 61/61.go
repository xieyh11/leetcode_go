package main

import (
	"../common"
)

type ListNode = common.ListNode

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	totalNum := 0
	temp := head
	for temp != nil {
		totalNum++
		temp = temp.Next
	}
	k %= totalNum
	if k > 0 {
		from := totalNum - k
		temp = head
		count := 0
		for temp != nil && count < from-1 {
			count++
			temp = temp.Next
		}
		curHead := temp.Next
		temp.Next = nil
		temp = curHead
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = head
		head = curHead
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head.PrintList()
	head = rotateRight(head, 6)
	head.PrintList()
}
