package main

import (
	"../common"
	"container/list"
	"fmt"
)

type TreeNode = common.TreeNode

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	iteList := list.New()
	iteList.PushFront(root)
	res := make([]int, 0)
	for iteList.Len() > 0 {
		front := iteList.Front()
		node := front.Value.(*TreeNode)
		res = append(res, node.Val)
		if node.Right != nil {
			iteList.PushFront(node.Right)
		}
		if node.Left != nil {
			iteList.PushFront(node.Left)
		}
		iteList.Remove(front)
	}
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(preorderTraversal(root))
}
