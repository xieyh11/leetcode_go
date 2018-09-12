package main

import (
	"../common"
)

type ListNode = common.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	current := head

	for current != nil {
		count := 0
		search := current
		searchEnd := search
		for ; searchEnd != nil && searchEnd.Val == search.Val; searchEnd = searchEnd.Next {
			count++
		}
		if count > 1 {
			current.Next = searchEnd
		}
		current = current.Next
	}
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	head.PrintList()
	head = deleteDuplicates(head)
	head.PrintList()
}
