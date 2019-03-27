package main

import (
	"../common"
	"fmt"
	"strconv"
)

type TreeNode = common.TreeNode

func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	fomerStr := strconv.FormatInt(int64(root.Val), 10)
	if root.Left == nil && root.Right == nil {
		return []string{fomerStr}
	}

	res := make([]string, 0)

	left := binaryTreePaths(root.Left)

	for _, path := range left {
		res = append(res, fomerStr+"->"+path)
	}

	right := binaryTreePaths(root.Right)
	for _, path := range right {
		res = append(res, fomerStr+"->"+path)
	}
	return res
}

func main() {
	root := common.ConstructTree([]string{"1", "2", "3", "null", "5"})
	fmt.Println(binaryTreePaths(root))
}
