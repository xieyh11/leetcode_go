package main

import (
	"../common"
)

type TreeNode = common.TreeNode
type ListNode = common.ListNode

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	twiceStep := head
	oneStep := head
	formerOneStep := (*ListNode)(nil)
	for twiceStep.Next != nil {
		formerOneStep = oneStep
		oneStep = oneStep.Next
		twiceStep = twiceStep.Next
		if twiceStep.Next == nil {
			break
		}
		twiceStep = twiceStep.Next
	}
	if formerOneStep != nil {
		formerOneStep.Next = nil
	} else {
		head = nil
	}
	return &TreeNode{Val: oneStep.Val, Left: sortedListToBST(head), Right: sortedListToBST(oneStep.Next)}
}

func main() {
	head := &ListNode{-10, &ListNode{-3, &ListNode{0, &ListNode{5, &ListNode{9, nil}}}}}
	root := sortedListToBST(head)
	root.Print()
}
