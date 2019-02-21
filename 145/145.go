package main

import (
	"../common"
	"container/list"
	"fmt"
)

type TreeNode = common.TreeNode

type NodeTraverse struct {
	Node      *TreeNode
	Traversed bool
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	iteList := list.New()
	iteList.PushFront(NodeTraverse{root, false})
	res := make([]int, 0)
	for iteList.Len() > 0 {
		front := iteList.Front()
		nodeTraverse := front.Value.(NodeTraverse)
		if nodeTraverse.Traversed {
			res = append(res, nodeTraverse.Node.Val)
			iteList.Remove(front)
		} else {
			if nodeTraverse.Node.Right != nil {
				iteList.PushFront(NodeTraverse{nodeTraverse.Node.Right, false})
			}
			if nodeTraverse.Node.Left != nil {
				iteList.PushFront(NodeTraverse{nodeTraverse.Node.Left, false})
			}
			nodeTraverse.Traversed = true
			front.Value = nodeTraverse
		}
	}
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(postorderTraversal(root))
}
