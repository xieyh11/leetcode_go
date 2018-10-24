package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func maxNode(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func minNode(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	okL := false
	okR := false
	if okL = isValidBST(left); okL {
		if leftMax := maxNode(left); leftMax != nil && leftMax.Val >= root.Val {
			return false
		}
	}
	if okR = isValidBST(right); okR {
		if rightMin := minNode(right); rightMin != nil && rightMin.Val <= root.Val {
			return false
		}
	}

	return okL && okR
}

func main() {
	root := &TreeNode{3, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}, &TreeNode{5, &TreeNode{4, nil, nil}, &TreeNode{6, nil, nil}}}
	root.Print()
	fmt.Println(isValidBST(root))
}
