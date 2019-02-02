package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func isBalanced(root *TreeNode) bool {
	temp, _ := isBalancedWithHeight(root)
	return temp
}

func isBalancedWithHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}
	leftOk, leftH := isBalancedWithHeight(root.Left)
	rightOk, rightH := isBalancedWithHeight(root.Right)
	maxH := leftH
	if maxH < rightH {
		maxH = rightH
	}
	if leftOk && rightOk {
		diff := leftH - rightH
		if diff >= -1 && diff <= 1 {
			return true, maxH + 1
		} else {
			return false, maxH + 1
		}
	} else {
		return false, maxH + 1
	}
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(isBalanced(root))
}
