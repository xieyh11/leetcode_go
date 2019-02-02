package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return sum == root.Val
	}
	remain := sum - root.Val
	left := false
	right := false
	if root.Left != nil {
		left = hasPathSum(root.Left, remain)
	}
	if root.Right != nil {
		right = hasPathSum(root.Right, remain)
	}
	return left || right
}

func main() {
	root := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, nil, &TreeNode{1, nil, nil}}}}
	fmt.Println(hasPathSum(root, 22))
}
