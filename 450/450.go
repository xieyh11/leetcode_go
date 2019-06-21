package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func findKey(root *TreeNode, key int) (*TreeNode, *TreeNode) {
	p := (*TreeNode)(nil)
	now := root
	for now != nil && now.Val != key {
		if now.Val > key {
			p, now = now, now.Left
		} else {
			p, now = now, now.Right
		}
	}
	return p, now
}

func rightLeftMost(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}
	p := root
	now := root.Right
	for now != nil && now.Left != nil {
		p, now = now, now.Left
	}
	return p, now
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	parent, node := findKey(root, key)
	if node == nil {
		return root
	}
	exP, exN := rightLeftMost(node)
	if exN == nil {
		if parent == nil {
			root = node.Left
		} else {
			if parent.Left == node {
				parent.Left = node.Left
			} else {
				parent.Right = node.Left
			}
		}
	} else {
		node.Val, exN.Val = exN.Val, node.Val
		if exP.Right == exN {
			exP.Right = exN.Right
		} else {
			exP.Left = exN.Right
		}
	}
	return root
}

func main() {
	treeStr := []string{"5", "3", "6", "2", "4", "null", "7"}
	root := common.ConstructTree(treeStr)
	key := 5
	root = deleteNode(root, key)
	root.Print()
}
