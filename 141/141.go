package main

import (
	"../common"
	"fmt"
)

type ListNode = common.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	if head.Next == nil {
		return false
	}
	stepOne := head
	stepTwo := head.Next.Next
	if stepTwo == nil {
		return false
	}
	for stepTwo != stepOne {
		if stepTwo.Next != nil && stepTwo.Next.Next != nil {
			stepOne = stepOne.Next
			stepTwo = stepTwo.Next.Next
		} else {
			break
		}
	}
	return stepOne == stepTwo
}

func main() {
	head := []*ListNode{&ListNode{3, nil}, &ListNode{2, nil}, &ListNode{0, nil}, &ListNode{-4, nil}}
	head[0].Next = head[1]
	head[1].Next = head[2]
	head[2].Next = head[3]
	head[3].Next = head[1]
	fmt.Println(hasCycle(head[0]))
}
