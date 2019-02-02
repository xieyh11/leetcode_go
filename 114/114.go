package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func walkRight(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	right := root.Right
	if root.Left != nil {
		flatten(root.Left)
		left := root.Left
		leftEnd := walkRight(left)
		root.Left = nil
		leftEnd.Right = right
		root.Right = left
	}
	if right != nil {
		flatten(root.Right)
	}
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}
	flatten(root)
	root.Print()
}
