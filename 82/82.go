package main

import (
	"../common"
)

type ListNode = common.ListNode

func deleteDuplicates(head *ListNode) *ListNode {
	tempHead := &ListNode{0, head}
	current := tempHead

	for current != nil {
		count := 0
		search := current.Next
		searchEnd := search
		for ; searchEnd != nil && searchEnd.Val == search.Val; searchEnd = searchEnd.Next {
			count++
		}
		if count == 1 {
			current = current.Next
		} else {
			current.Next = searchEnd
		}
		if count == 0 {
			break
		}
	}
	return tempHead.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}}
	head.PrintList()
	head = deleteDuplicates(head)
	head.PrintList()
}
