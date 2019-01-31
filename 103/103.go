package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func zigzagLevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	currentNode := make([]*TreeNode, 1)
	nextNode := make([]*TreeNode, 0)

	if root == nil {
		return res
	}

	currentNode[0] = root
	reverse := false

	for currentNode != nil && len(currentNode) > 0 {
		tempRes := make([]int, 0)
		for i := len(currentNode) - 1; i >= 0; i-- {
			tempRes = append(tempRes, currentNode[i].Val)
			if reverse {
				if currentNode[i].Right != nil {
					nextNode = append(nextNode, currentNode[i].Right)
				}
				if currentNode[i].Left != nil {
					nextNode = append(nextNode, currentNode[i].Left)
				}
			} else {
				if currentNode[i].Left != nil {
					nextNode = append(nextNode, currentNode[i].Left)
				}
				if currentNode[i].Right != nil {
					nextNode = append(nextNode, currentNode[i].Right)
				}
			}
		}
		reverse = !reverse
		res = append(res, tempRes)
		currentNode = nextNode
		nextNode = make([]*TreeNode, 0)
	}
	return res
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(zigzagLevelOrder(root))
}
