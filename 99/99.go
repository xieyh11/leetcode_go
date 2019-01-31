package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func getRightMost(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Right != nil {
		root = root.Right
	}
	return root
}

func getLeftMost(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func isBinarySearchTree(root *TreeNode) (bool, []*TreeNode) {
	if root == nil {
		return true, []*TreeNode{}
	}
	left := root.Left
	right := root.Right
	okL, errL := isBinarySearchTree(left)
	okR, errR := isBinarySearchTree(right)
	if len(errL) == 1 {
		temp := errL[0]
		if temp.Left != nil && temp.Left.Val > temp.Val {
			errL[0] = temp.Left
		}
	}
	if len(errR) == 1 {
		temp := errR[0]
		if temp.Right != nil && temp.Right.Val < temp.Val {
			errR[0] = temp.Right
		}
	}
	switch {
	case okL && okR:
		res := make([]*TreeNode, 0)
		if left != nil {
			leftMax := getRightMost(left)
			if leftMax.Val > root.Val {
				res = append(res, leftMax)
			}
		}
		if right != nil {
			rightMin := getLeftMost(right)
			if rightMin.Val < root.Val {
				res = append(res, rightMin)
			}
		}
		return len(res) == 0, res
	case !okL && okR:
		if len(errL) == 2 {
			return okL, errL
		} else {
			if right != nil {
				rightMin := getLeftMost(right)
				if rightMin.Val < root.Val {
					errL = append(errL, rightMin)
					return false, errL
				}
			}
			if errL[0].Val > root.Val {

			}
			return false, errL
		}
	case okL && !okR:
		if len(errR) == 2 {
			return okR, errR
		} else {
			if left != nil {
				leftMax := getRightMost(left)
				if leftMax.Val > root.Val {
					errR = append(errR, leftMax)
					return false, errR
				}
			}
			errR = append
		}
	default:
		return false, root
	}
}

func recoverTree(root *TreeNode) {
	if root == nil {
		return
	}

}

func main() {
	// root := &TreeNode{1, &TreeNode{3, nil, &TreeNode{2, nil, nil}}, nil}
	// root := &TreeNode{3, &TreeNode{1, nil, nil}, &TreeNode{4, &TreeNode{2, nil, nil}, nil}}
	root := &TreeNode{146, &TreeNode{71, &TreeNode{55, &TreeNode{321, &TreeNode{-33, nil, nil}, nil}, nil}, nil}, &TreeNode{-13, &TreeNode{231, nil, nil}, &TreeNode{399, nil, nil}}}
	root.Print()
	recoverTree(root)
	root.Print()
}
