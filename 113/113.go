package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return [][]int{[]int{sum}}
		}
	}
	remain := sum - root.Val
	left := [][]int{}
	right := [][]int{}
	if root.Left != nil {
		left = pathSum(root.Left, remain)
	}
	if root.Right != nil {
		right = pathSum(root.Right, remain)
	}
	total := append(left, right...)
	if len(total) > 0 {
		rootInt := []int{root.Val}
		for i := range total {
			total[i] = append(rootInt, total[i]...)
		}
	}
	return total
}

func main() {
	root := &TreeNode{5, &TreeNode{4, &TreeNode{11, &TreeNode{7, nil, nil}, &TreeNode{2, nil, nil}}, nil}, &TreeNode{8, &TreeNode{13, nil, nil}, &TreeNode{4, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}}}}
	fmt.Println(pathSum(root, 22))
}
