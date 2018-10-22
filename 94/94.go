package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	left := root.Left
	right := root.Right
	res := []int{}
	if left != nil {
		res = append(res, inorderTraversal(left)...)
	}
	res = append(res, root.Val)
	if right != nil {
		res = append(res, inorderTraversal(right)...)
	}
	return res
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{2, &TreeNode{3, nil, nil}, nil}}
	fmt.Println(inorderTraversal(root))
}
