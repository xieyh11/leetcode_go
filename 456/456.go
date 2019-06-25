package main

import (
	"fmt"
)

type AVLTree struct {
	Root *AVLTreeNode
}

type AVLTreeNode struct {
	Val     int
	Parent  *AVLTreeNode
	Left    *AVLTreeNode
	Right   *AVLTreeNode
	Balance int
}

func RotateLeftRight(parent, node *AVLTreeNode) *AVLTreeNode {
	Y := node.Right
	t3 := Y.Left
	node.Right = t3
	if t3 != nil {
		t3.Parent = node
	}
	Y.Left = node
	node.Parent = Y
	t2 := Y.Right
	parent.Left = t2
	if t2 != nil {
		t2.Parent = parent
	}
	Y.Right = parent
	parent.Parent = Y
	if Y.Balance > 0 {
		parent.Balance = -1
		node.Balance = 0
	} else if Y.Balance == 0 {
		parent.Balance = 0
		node.Balance = 0
	} else {
		parent.Balance = 0
		node.Balance = 1
	}
	Y.Balance = 0
	return Y
}

func RotateRightLeft(parent, node *AVLTreeNode) *AVLTreeNode {
	Y := node.Left
	t3 := Y.Right
	node.Left = t3
	if t3 != nil {
		t3.Parent = node
	}
	Y.Right = node
	node.Parent = Y
	t2 := Y.Left
	parent.Right = t2
	if t2 != nil {
		t2.Parent = parent
	}
	Y.Left = parent
	parent.Parent = Y
	if Y.Balance > 0 {
		parent.Balance = -1
		node.Balance = 0
	} else if Y.Balance == 0 {
		parent.Balance = 0
		node.Balance = 0
	} else {
		parent.Balance = 0
		node.Balance = 1
	}
	Y.Balance = 0
	return Y
}

func RotateLeft(parent, node *AVLTreeNode) *AVLTreeNode {
	t23 := node.Left
	parent.Right = t23
	if t23 != nil {
		t23.Parent = parent
	}
	node.Left = parent
	parent.Parent = node

	if node.Balance == 0 {
		parent.Balance = 1
		node.Balance = -1
	} else {
		parent.Balance = 0
		node.Balance = 0
	}
	return node
}

func RotateRight(parent, node *AVLTreeNode) *AVLTreeNode {
	t23 := node.Right
	parent.Left = t23
	if t23 != nil {
		t23.Parent = parent
	}
	node.Right = parent
	parent.Parent = node

	if node.Balance == 0 {
		parent.Balance = 1
		node.Balance = -1
	} else {
		parent.Balance = 0
		node.Balance = 0
	}
	return node
}

func InsertFind(root *AVLTreeNode, val int) (*AVLTreeNode, bool, bool) {
	if root.Val == val {
		return root, false, false
	}
	if root.Val > val {
		if root.Left == nil {
			return root, true, false
		}
		return InsertFind(root.Left, val)
	} else {
		if root.Right == nil {
			return root, false, true
		}
		return InsertFind(root.Right, val)
	}
}

func (this *AVLTree) Insert(val int) {
	if this.Root == nil {
		this.Root = &AVLTreeNode{Val: val}
		return
	}
	node, onLeft, onRight := InsertFind(this.Root, val)
	if !onLeft && !onRight {
		return
	}
	parent := node
	if onLeft {
		node.Left = &AVLTreeNode{Val: val, Parent: node}
		node = node.Left
	}
	if onRight {
		node.Right = &AVLTreeNode{Val: val, Parent: node}
		node = node.Right
	}
	for parent != nil {
		G := parent.Parent
		N := (*AVLTreeNode)(nil)
		if node == parent.Left {
			if parent.Balance < 0 {
				if node.Balance > 0 {
					N = RotateLeftRight(parent, node)
				} else {
					N = RotateRight(parent, node)
				}
			} else {
				if parent.Balance > 0 {
					parent.Balance = 0
					break
				}
				parent.Balance = -1
				node = parent
				parent = node.Parent
				continue
			}
		} else {
			if parent.Balance > 0 {
				if node.Balance < 0 {
					N = RotateRightLeft(parent, node)
				} else {
					N = RotateLeft(parent, node)
				}
			} else {
				if parent.Balance < 0 {
					parent.Balance = 0
					break
				}
				parent.Balance = 1
				node, parent = parent, parent.Parent
				continue
			}
		}
		N.Parent = G
		if G != nil {
			if parent == G.Left {
				G.Left = N
			} else {
				G.Right = N
			}
		} else {
			this.Root = N
		}
		break
	}
}

func Constructor() *AVLTree {
	return &AVLTree{}
}

func (this *AVLTree) findBetween(s, l int) bool {
	root := this.Root
	for root != nil {
		if root.Val > s && root.Val < l {
			return true
		} else if root.Val <= s {
			root = root.Right
		} else {
			root = root.Left
		}
	}
	return false
}

func find132pattern(nums []int) bool {
	n := len(nums)
	if n < 3 {
		return false
	}

	minStack := make([]int, 1)
	minStack[0] = 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[minStack[len(minStack)-1]] {
			minStack = append(minStack, i)
		}
	}
	stackLast := len(minStack) - 1
	tree := Constructor()
	for i := n - 1; i > 0; i-- {
		if minStack[stackLast] >= i {
			stackLast--
		}
		if tree.findBetween(nums[minStack[stackLast]], nums[i]) {
			return true
		}
		tree.Insert(nums[i])
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(find132pattern(nums))

	nums = []int{3, 1, 4, 2}
	fmt.Println(find132pattern(nums))

	nums = []int{-1, 3, 2, 0}
	fmt.Println(find132pattern(nums))
}
