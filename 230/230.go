package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func recSol(root *TreeNode, k int) (int, bool, int) {
	if root == nil {
		return 0, false, 0
	}
	left, leftFound, leftCount := recSol(root.Left, k)
	if leftFound {
		return left, leftFound, 0
	}
	if leftCount+1 == k {
		return root.Val, true, 0
	}
	right, rightFound, rightCount := recSol(root.Right, k-leftCount-1)
	if rightFound {
		return right, rightFound, 0
	}
	return root.Val, false, leftCount + 1 + rightCount
}

func kthSmallest(root *TreeNode, k int) int {
	num, _, _ := recSol(root, k)
	return num
}

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}
	for k := 1; k <= 6; k++ {
		fmt.Println(kthSmallest(root, k))
	}
}
