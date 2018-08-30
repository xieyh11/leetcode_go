package common

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Print the single-linked list.
func (head *ListNode) PrintList() {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print("->")
		}
		head = head.Next
	}
	fmt.Println()
}

// Definition for an interval.
type Interval struct {
	Start int
	End   int
}
