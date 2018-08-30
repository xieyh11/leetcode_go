package main

import (
	"../common"
)

type ListNode = common.ListNode

func reverseKGroup(head *ListNode, k int) *ListNode {
	fakeHead := &ListNode{0, head}
	tempAt := fakeHead
	storeK := make([]*ListNode, k)
	for true {
		count := 0
		countAt := tempAt
		for countAt.Next != nil && count < k {
			storeK[count] = countAt.Next
			countAt = countAt.Next
			count++
		}
		if count == k {
			lastK := countAt.Next
			for i := range storeK {
				if i == 0 {
					storeK[i].Next = lastK
				} else {
					storeK[i].Next = storeK[i-1]
				}
			}
			tempAt.Next = storeK[k-1]
			tempAt = storeK[0]
		} else {
			break
		}
	}
	return fakeHead.Next
}

func main() {
	head := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	head.PrintList()

	head = reverseKGroup(head, 2)
	head.PrintList()
}
