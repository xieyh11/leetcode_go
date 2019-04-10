package main

import (
	"fmt"
)

type AVLTree struct {
	Root *AVLTreeNode
}

type AVLTreeNode struct {
	Val          int
	Count        int
	Parent       *AVLTreeNode
	Left         *AVLTreeNode
	Right        *AVLTreeNode
	Balance      int
	SubtreeCount int
}

func InsertFind(root *AVLTreeNode, val int) (*AVLTreeNode, bool, bool) {
	root.SubtreeCount++
	if root.Val == val {
		root.Count++
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

func UpdateCount(node *AVLTreeNode) {
	node.SubtreeCount = node.Count
	if node.Left != nil {
		node.SubtreeCount += node.Left.SubtreeCount
	}
	if node.Right != nil {
		node.SubtreeCount += node.Right.SubtreeCount
	}
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
	UpdateCount(node)
	UpdateCount(parent)
	UpdateCount(Y)
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
	UpdateCount(node)
	UpdateCount(parent)
	UpdateCount(Y)
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
	UpdateCount(parent)
	UpdateCount(node)
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
	UpdateCount(parent)
	UpdateCount(node)
	return node
}

func (this *AVLTree) Insert(val int) {
	if this.Root == nil {
		this.Root = &AVLTreeNode{Val: val, Count: 1, SubtreeCount: 1}
		return
	}
	node, onLeft, onRight := InsertFind(this.Root, val)
	if !onLeft && !onRight {
		return
	}
	parent := node
	if onLeft {
		node.Left = &AVLTreeNode{Val: val, Count: 1, SubtreeCount: 1, Parent: node}
		node = node.Left
	}
	if onRight {
		node.Right = &AVLTreeNode{Val: val, Count: 1, SubtreeCount: 1, Parent: node}
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

func (this *AVLTree) CountSmaller(val int) int {
	if this.Root == nil {
		return 0
	}
	now := this.Root
	res := 0
	for now != nil {
		if now.Val > val {
			now = now.Left
		} else {
			if now.Left != nil {
				res += now.Left.SubtreeCount
			}
			if now.Val < val {
				res += now.Count
				now = now.Right
			} else {
				break
			}
		}
	}
	return res
}

func countSmaller(nums []int) []int {
	tree := Constructor()
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = tree.CountSmaller(nums[i])
		tree.Insert(nums[i])
	}
	return res
}

func main() {
	nums := []int{26, 78, 27, 100, 33, 67, 90, 23, 66, 5, 38, 7, 35, 23, 52, 22, 83, 51, 98, 69, 81, 32, 78, 28, 94, 13, 2, 97, 3, 76, 99, 51, 9, 21, 84, 66, 65, 36, 100, 41}
	fmt.Println(countSmaller(nums))
}
