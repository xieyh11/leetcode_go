package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	count := 0
	for i := head; i != nil; i = i.Next {
		count++
	}
	if count < n {
		return head
	}
	idxB := count - n + 1
	if idxB == 1 {
		head = head.Next
	} else {
		i := head.Next
		bi := head
		for idxB > 2 {
			bi = i
			i = bi.Next
			idxB--
		}
		bi.Next = i.Next
	}
	return head
}

func printList(head *ListNode) {
	for ; head != nil; head = head.Next {
		fmt.Println(head.Val)
	}
}

func main() {
	head := new(ListNode)
	head.Val = 1
	for i, j := 2, head; i < 6; i++ {
		j.Next = new(ListNode)
		j = j.Next
		j.Val = i
	}
	printList(head)
	fmt.Println()
	head = removeNthFromEnd(head, 2)
	printList(head)
}
