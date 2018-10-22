package main

import (
	"../common"
)

type ListNode = common.ListNode

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	first := head
	if first == nil {
		return head, head
	}
	second := head.Next
	if second == nil {
		return head, head
	}
	tail := head
	formerNode := (*ListNode)(nil)
	for second != nil {
		nextNode := second.Next
		if nextNode == nil {
			head = second
		}
		second.Next = first
		first.Next = formerNode
		formerNode, first, second = first, second, nextNode
	}
	return head, tail
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	count := 0
	var formerNode *ListNode
	if m == 1 {
		formerNode = (*ListNode)(nil)
	} else {
		nowAt := head
		for nowAt != nil {
			count++
			formerNode, nowAt = nowAt, nowAt.Next
			if count >= m-1 {
				break
			}
		}
	}
	mFormer := formerNode
	var nNext *ListNode
	var fakeHead *ListNode
	if formerNode == nil {
		fakeHead = head
	} else {
		fakeHead = formerNode.Next
	}
	first := fakeHead
	if first == nil {
		return head
	}
	if m == n {
		return head
	}
	second := first.Next
	if second == nil {
		return head
	}
	m++
	fakeTail := fakeHead
	for m <= n {
		nextNode := second.Next
		if m == n {
			fakeHead = second
			nNext = nextNode
		}
		second.Next = first
		first.Next = formerNode
		formerNode, first, second = first, second, nextNode
		m++
	}
	if mFormer == nil {
		head = fakeHead
	} else {
		mFormer.Next = fakeHead
	}
	fakeTail.Next = nNext
	return head
}

func main() {
	head := &ListNode{1, &ListNode{2, nil}}
	head.PrintList()
	head = reverseBetween(head, 1, 1)
	head.PrintList()
}
