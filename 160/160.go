package main

import (
	"../common"
	"fmt"
)

type ListNode = common.ListNode

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return nil
	}
	stepOne := head
	stepTwo := head
	for true {
		if stepTwo.Next != nil && stepTwo.Next.Next != nil {
			stepOne = stepOne.Next
			stepTwo = stepTwo.Next.Next
		} else {
			return nil
		}
		if stepOne == stepTwo {
			break
		}
	}
	stepOne = head
	for stepOne != stepTwo {
		stepOne = stepOne.Next
		stepTwo = stepTwo.Next
	}
	return stepOne
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	lastB := headB
	for lastB.Next != nil {
		lastB = lastB.Next
	}
	lastB.Next = headA
	defer func(node *ListNode) { node.Next = nil }(lastB)

	return detectCycle(headB)
}

func main() {
	common := &ListNode{8, &ListNode{4, &ListNode{5, nil}}}
	headA := &ListNode{4, &ListNode{1, common}}
	headB := &ListNode{5, &ListNode{0, &ListNode{1, common}}}
	res := getIntersectionNode(headA, headB)
	fmt.Println(res.Val)
}
