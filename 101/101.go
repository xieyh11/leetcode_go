package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func isSymmetricRec(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}

	if left.Val == right.Val {
		return isSymmetricRec(left.Left, right.Right) && isSymmetricRec(left.Right, right.Left)
	} else {
		return false
	}
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricRec(root.Left, root.Right)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	fmt.Println(isSymmetric(root))
}
