package main

import (
	"../common"
)

type ListNode = common.ListNode

func deleteNode(node *ListNode) {
	if node == nil || node.Next == nil {
		return
	}
	former, now := node, node.Next
	former.Val = now.Val
	for now.Next != nil {
		former, now = now, now.Next
		former.Val = now.Val
	}
	former.Next = nil
}

func main() {
	node := &ListNode{3, &ListNode{4, nil}}
	head := &ListNode{1, &ListNode{2, node}}
	head.PrintList()
	deleteNode(node)
	head.PrintList()
}
