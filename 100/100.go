package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	if p.Val == q.Val {
		return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	} else {
		return false
	}
}

func main() {
	p := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	q := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(isSameTree(p, q))
}
