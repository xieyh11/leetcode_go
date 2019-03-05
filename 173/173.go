package main

import (
	"../common"
	"container/list"
	"fmt"
)

type TreeNode = common.TreeNode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type BSTIterator struct {
	Heap *list.List
}

func Constructor(root *TreeNode) BSTIterator {
	nowAt := root
	heap := list.New()
	for nowAt != nil {
		heap.PushBack(nowAt)
		nowAt = nowAt.Left
	}
	return BSTIterator{heap}
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	back := this.Heap.Remove(this.Heap.Back()).(*TreeNode)
	nowAt := back.Right
	for nowAt != nil {
		this.Heap.PushBack(nowAt)
		nowAt = nowAt.Left
	}
	return back.Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.Heap.Len() > 0
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {
	root := &TreeNode{7, &TreeNode{3, nil, nil}, &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	iterator := Constructor(root)
	fmt.Println(iterator.Next())    // return 3
	fmt.Println(iterator.Next())    // return 7
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 9
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 15
	fmt.Println(iterator.HasNext()) // return true
	fmt.Println(iterator.Next())    // return 20
	fmt.Println(iterator.HasNext()) // return false
}
