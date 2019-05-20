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

// Construct tree from an array
func ConstructTree(arr []string) *TreeNode {
	if len(arr) == 0 || arr[0] == "null" {
		return nil
	}
	root := &TreeNode{}
	temp, _ := strconv.ParseInt(arr[0], 10, 64)
	root.Val = int(temp)

	search := list.New()
	search.PushBack(root)
	nowAt := 1
	for nowAt < len(arr) && search.Len() > 0 {
		node := search.Remove(search.Front()).(*TreeNode)
		if arr[nowAt] != "null" {
			temp, _ := strconv.ParseInt(arr[nowAt], 10, 64)
			node.Left = &TreeNode{Val: int(temp)}
			search.PushBack(node.Left)
		}
		nowAt++
		if nowAt < len(arr) && arr[nowAt] != "null" {
			temp, _ := strconv.ParseInt(arr[nowAt], 10, 64)
			node.Right = &TreeNode{Val: int(temp)}
			search.PushBack(node.Right)
		}
		nowAt++
	}
	return root
}

// Definition for a point.
type Point struct {
	X int
	Y int
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

var Bits uint = 32 << (^uint(0) >> 63)
