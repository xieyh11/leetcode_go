package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func isLeaf(root *TreeNode) bool {
	return root.Left == nil && root.Right == nil
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	if root.Left != nil {
		if isLeaf(root.Left) {
			res += root.Left.Val
		} else {
			res += sumOfLeftLeaves(root.Left)
		}
	}
	if root.Right != nil {
		if !isLeaf(root.Right) {
			res += sumOfLeftLeaves(root.Right)
		}
	}
	return res
}

func main() {
	root := common.ConstructTree([]string{"3", "9", "20", "null", "null", "15", "7"})
	fmt.Println(sumOfLeftLeaves(root))
}
