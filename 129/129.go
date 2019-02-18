package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func recSol(root *TreeNode, parent int) int {
	childParent := parent*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return childParent
	}
	sum := 0
	if root.Left != nil {
		sum += recSol(root.Left, childParent)
	}
	if root.Right != nil {
		sum += recSol(root.Right, childParent)
	}
	return sum
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recSol(root, 0)
}

func main() {
	root := &TreeNode{4, &TreeNode{9, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}, &TreeNode{0, nil, nil}}
	fmt.Println(sumNumbers(root))
}
