package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	currentDepth := 1
	if root.Left != nil {
		currentDepth = 1 + maxDepth(root.Left)
	}

	if root.Right != nil {
		rightDepth := 1 + maxDepth(root.Right)
		if rightDepth > currentDepth {
			currentDepth = rightDepth
		}
	}
	return currentDepth
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(maxDepth(root))
}
