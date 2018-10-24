package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func inodeSearchAdding(root *TreeNode, addOn int) {
	if root == nil {
		return
	}

	left := root.Left
	right := root.Right
	inodeSearchAdding(left, addOn)
	root.Val += addOn
	inodeSearchAdding(right, addOn)
	return
}

func inodeCopy(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	left := root.Left
	right := root.Right
	res := new(TreeNode)
	res.Val = root.Val
	res.Left = inodeCopy(left)
	res.Right = inodeCopy(right)
	return res
}

func inodeCopyAdd(root *TreeNode, addOn int) *TreeNode {
	if root == nil {
		return root
	}
	left := root.Left
	right := root.Right
	res := new(TreeNode)
	res.Val = root.Val + addOn
	res.Left = inodeCopyAdd(left, addOn)
	res.Right = inodeCopyAdd(right, addOn)
	return res
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	res := make([][]*TreeNode, n+1)
	res[0] = []*TreeNode{nil}
	res[1] = []*TreeNode{&TreeNode{1, nil, nil}}

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			for _, left := range res[j] {
				for _, right := range res[i-1-j] {
					tempR := new(TreeNode)
					tempR.Val = j + 1
					tempR.Left = left
					tempR.Right = inodeCopyAdd(right, tempR.Val)
					res[i] = append(res[i], tempR)
				}
			}
		}
	}
	return res[n]
}

func main() {
	res := generateTrees(3)
	for _, v := range res {
		v.Print()
	}
}
