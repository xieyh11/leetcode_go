package main

import (
	"../common"
	"fmt"
)

type ListNode = common.ListNode

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	former, nowat := head, head.Next
	head.Next = nil
	for nowat != nil {
		temp := nowat.Next
		nowat.Next = former
		former, nowat = nowat, temp
	}
	return former
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	stepOne := head
	stepTwo := head
	for stepTwo.Next != nil && stepTwo.Next.Next != nil {
		stepOne = stepOne.Next
		stepTwo = stepTwo.Next.Next
	}
	halfHead := stepOne.Next
	stepOne.Next = nil
	halfHead = reverseList(halfHead)
	for first, second := head, halfHead; first != nil && second != nil; first, second = first.Next, second.Next {
		if first.Val != second.Val {
			return false
		}
	}
	return true
}

func main() {
	head := &ListNode{1, &ListNode{3, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}
	fmt.Println(isPalindrome(head))
}
