package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := invertTree(root.Left)
	right := invertTree(root.Right)
	root.Left, root.Right = right, left
	return root
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}
	root.Print()
	root = invertTree(root)
	root.Print()
}
