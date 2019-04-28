package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func recSol(root *TreeNode, store map[*TreeNode]int) int {
	if root == nil {
		return 0
	}
	if v, ok := store[root]; ok {
		return v
	}
	oneRes := recSol(root.Left, store) + recSol(root.Right, store)
	anotherRes := root.Val
	if root.Left != nil {
		anotherRes += recSol(root.Left.Left, store) + recSol(root.Left.Right, store)
	}
	if root.Right != nil {
		anotherRes += recSol(root.Right.Left, store) + recSol(root.Right.Right, store)
	}
	if oneRes < anotherRes {
		oneRes = anotherRes
	}
	store[root] = oneRes
	return oneRes
}

func rob(root *TreeNode) int {
	store := make(map[*TreeNode]int)
	return recSol(root, store)
}

func main() {
	root0 := []string{"3", "2", "3", "null", "3", "null", "1"}
	fmt.Println(rob(common.ConstructTree(root0)) == 7)

	root1 := []string{"3", "4", "5", "1", "3", "null", "1"}
	fmt.Println(rob(common.ConstructTree(root1)) == 9)
}
