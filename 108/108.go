package main

import (
	"../common"
)

type TreeNode = common.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	rootIdx := len(nums) / 2
	root := new(TreeNode)
	root.Val = nums[rootIdx]
	root.Left = sortedArrayToBST(nums[:rootIdx])
	root.Right = sortedArrayToBST(nums[rootIdx+1:])
	return root
}

func main() {
	nums := []int{-10, -3, 0, 5, 9}
	root := sortedArrayToBST(nums)
	root.Print()
}
