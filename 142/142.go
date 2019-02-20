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

func main() {
	headInt := []int{3, 2, 0, -4}
	pos := 1
	head := make([]*ListNode, len(headInt))
	for i := len(head) - 1; i >= 0; i-- {
		next := (*ListNode)(nil)
		if i != len(head)-1 {
			next = head[i+1]
		}
		head[i] = &ListNode{headInt[i], next}
	}
	head[len(head)-1].Next = head[pos]
	fmt.Println(headInt[pos], detectCycle(head[0]).Val)
}
