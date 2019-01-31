package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = preorder[0]
	rootIdx := 0
	for inorder[rootIdx] != preorder[0] {
		rootIdx++
	}
	leftInorder := inorder[:rootIdx]
	rightInorder := inorder[rootIdx+1:]
	leftPreorder := preorder[1 : 1+len(leftInorder)]
	rightPreorder := preorder[1+len(leftInorder):]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(rightPreorder, rightInorder)
	return root
}

func main() {
	pre := []int{3, 9, 20, 15, 7}
	ino := []int{9, 3, 15, 20, 7}
	root := buildTree(pre, ino)
	root.Print()
}
