package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	rootIdx := 0
	for inorder[rootIdx] != root.Val {
		rootIdx++
	}
	leftInorder := inorder[:rootIdx]
	rightInorder := inorder[rootIdx+1:]
	leftPostorder := postorder[0:len(leftInorder)]
	rightPostorder := postorder[len(leftInorder) : len(postorder)-1]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}

func main() {
	pos := []int{9, 15, 7, 20, 3}
	ino := []int{9, 3, 15, 20, 7}
	root := buildTree(ino, pos)
	root.Print()
}
