package main

import (
	"../common"
)

type ListNode = common.ListNode

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nowAt := head
	next := head.Next
	head.Next = nil
	for next != nil {
		next.Next, nowAt, next = nowAt, next, next.Next
	}
	return nowAt
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1R := reverse(l1)
	l2R := reverse(l2)
	res := (*ListNode)(nil)
	nowAt := res
	carry := 0
	for l1R != nil && l2R != nil {
		temp := l1R.Val + l2R.Val + carry
		carry = temp / 10
		temp %= 10
		if res == nil {
			res = &ListNode{temp, nil}
			nowAt = res
		} else {
			nowAt.Next = &ListNode{temp, nil}
			nowAt = nowAt.Next
		}
		l1R = l1R.Next
		l2R = l2R.Next
	}
	if l1R == nil {
		l1R = l2R
	}
	for l1R != nil {
		temp := l1R.Val + carry
		carry = temp / 10
		temp %= 10
		if res == nil {
			res = &ListNode{temp, nil}
			nowAt = res
		} else {
			nowAt.Next = &ListNode{temp, nil}
			nowAt = nowAt.Next
		}
		l1R = l1R.Next
	}
	if carry != 0 {
		nowAt.Next = &ListNode{carry, nil}
	}
	return reverse(res)
}

func main() {
	l1 := &ListNode{7, &ListNode{2, &ListNode{4, &ListNode{3, nil}}}}
	l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	res := addTwoNumbers(l1, l2)
	res.PrintList()
}
