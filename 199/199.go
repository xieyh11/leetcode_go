package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	leftSide := rightSideView(root.Left)
	rightSide := rightSideView(root.Right)
	if len(rightSide) < len(leftSide) {
		rightSide = append(rightSide, leftSide[len(rightSide):]...)
	}
	return append([]int{root.Val}, rightSide...)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(rightSideView(root))
}
