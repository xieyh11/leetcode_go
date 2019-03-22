package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	if p == q {
		return p
	}
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func main() {
	p, q := &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}
	root := &TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, p, q}}, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
