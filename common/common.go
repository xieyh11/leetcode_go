package common

import (
	"container/list"
	"fmt"
	"strconv"
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

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (root *TreeNode) Print() {
	res := []string{}
	l := list.New()
	l.PushBack(root)
	for e := l.Front(); e != nil; e = e.Next() {
		nowPtr := e.Value.(*TreeNode)
		if nowPtr == nil {
			res = append(res, "null")
		} else {
			res = append(res, strconv.FormatInt(int64(nowPtr.Val), 10))
			l.PushBack(nowPtr.Left)
			l.PushBack(nowPtr.Right)
		}
	}
	fmt.Println(res)
}

// Definition for a point.
type Point struct {
	X int
	Y int
}
