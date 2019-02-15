package main

import (
	"../common"
	"fmt"
)

type TreeNode = common.TreeNode

func resSol(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	maxPath := root.Val

	if root.Left != nil && root.Right != nil {
		leftR, leftP := resSol(root.Left)
		if maxPath < leftP {
			maxPath = leftP
		}
		rightR, rightP := resSol(root.Right)
		if maxPath < rightP {
			maxPath = rightP
		}
		subMax := leftR
		if subMax < rightR {
			subMax = rightR
		}
		rootR := root.Val
		if subMax > 0 {
			rootR += subMax
		}

		if maxPath < rootR {
			maxPath = rootR
		}

		if maxPath < leftR+root.Val+rightR {
			maxPath = leftR + root.Val + rightR
		}
		return rootR, maxPath
	} else if root.Left == nil && root.Right == nil {
		return root.Val, root.Val
	} else if root.Left == nil {
		rightR, rightP := resSol(root.Right)
		if maxPath < rightP {
			maxPath = rightP
		}
		rootR := root.Val
		if rightR > 0 {
			rootR += rightR
		}
		if maxPath < rootR {
			maxPath = rootR
		}
		return rootR, maxPath
	} else {
		leftR, leftP := resSol(root.Left)
		if maxPath < leftP {
			maxPath = leftP
		}
		rootR := root.Val
		if leftR > 0 {
			rootR += leftR
		}
		if maxPath < rootR {
			maxPath = rootR
		}
		return rootR, maxPath
	}

}

func maxPathSum(root *TreeNode) int {
	_, res := resSol(root)
	return res
}

func main() {
	root := &TreeNode{-10, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(maxPathSum(root))
}
