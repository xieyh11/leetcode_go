package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func levelOrderBottom(root *TreeNode) [][]int {
	res := make([][]int, 0)
	currentNode := make([]*TreeNode, 1)
	nextNode := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	currentNode[0] = root

	for currentNode != nil && len(currentNode) > 0 {
		tempRes := make([]int, 0)
		for _, i := range currentNode {
			tempRes = append(tempRes, i.Val)
			if i.Left != nil {
				nextNode = append(nextNode, i.Left)
			}
			if i.Right != nil {
				nextNode = append(nextNode, i.Right)
			}
		}
		res = append([][]int{tempRes}, res...)
		currentNode = nextNode
		nextNode = make([]*TreeNode, 0)
	}
	return res
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(levelOrderBottom(root))
}
