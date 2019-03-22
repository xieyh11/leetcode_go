package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func recSol(root, p, q *TreeNode) (*TreeNode, bool, bool) {
	if root == nil {
		return root, false, false
	}
	left, leftp, leftq := recSol(root.Left, p, q)
	if left == nil {
		if (leftp && root == q) || (leftq && root == p) {
			return root, true, true
		}
		right, rightp, rightq := recSol(root.Right, p, q)
		if right == nil {
			if (rightp && root == q) || (rightq && root == p) {
				return root, true, true
			}
			if (leftp && rightq) || (leftq && rightp) {
				return root, true, true
			}
			return nil, (leftp || rightp || root == p), (leftq || rightq || root == q)
		}
		return right, rightp, rightq
	}
	return left, leftp, leftq
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	res, _, _ := recSol(root, p, q)
	return res
}

func main() {
	p, q := &TreeNode{3, nil, nil}, &TreeNode{8, nil, nil}
	root := &TreeNode{-1, &TreeNode{0, &TreeNode{-2, q, nil}, &TreeNode{4, nil, nil}}, p}
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
