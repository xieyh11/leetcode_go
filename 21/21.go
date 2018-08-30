package main

import (
	"../common"
)

type ListNode = common.ListNode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode = nil
	tempHead := head
	for l1 != nil && l2 != nil {
		temp := l1
		if temp.Val > l2.Val {
			temp = l2
			l2 = l2.Next
		} else {
			l1 = l1.Next
		}
		if head == nil {
			head = temp
			tempHead = head
		} else {
			tempHead.Next = temp
			tempHead = tempHead.Next
		}
	}
	if l1 != nil {
		if head == nil {
			head = l1
		} else {
			tempHead.Next = l1
		}
	}
	if l2 != nil {
		if head == nil {
			head = l2
		} else {
			tempHead.Next = l2
		}
	}
	return head
}

func main() {
	l1, l2 := [3]ListNode{}, [3]ListNode{}
	l1[0].Val, l1[0].Next = 1, &l1[1]
	l1[1].Val, l1[1].Next = 2, &l1[2]
	l1[2].Val, l1[2].Next = 4, nil

	l2[0].Val, l2[0].Next = 1, &l2[1]
	l2[1].Val, l2[1].Next = 3, &l2[2]
	l2[2].Val, l2[2].Next = 4, nil

	l1[0].PrintList()
	l2[0].PrintList()

	head := mergeTwoLists(&l1[0], &l2[0])
	head.PrintList()
}
