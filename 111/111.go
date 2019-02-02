package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := minDepth(root.Left)
	right := minDepth(root.Right)
	if left == 0 {
		return right + 1
	}
	if right == 0 {
		return left + 1
	}
	if left < right {
		return left + 1
	} else {
		return right + 1
	}
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(minDepth(root))
}
