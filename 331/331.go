package main

import (
	"container/list"
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	splits := strings.Split(preorder, ",")
	if len(splits) == 0 {
		return false
	}
	if splits[0] == "#" {
		return len(splits) == 1
	}
	stack := list.New()
	stack.PushBack(0)
	stack.PushBack(0)
	for i := 1; i < len(splits); i++ {
		if stack.Len() == 0 {
			return false
		}
		stack.Remove(stack.Back())
		if splits[i] != "#" {
			stack.PushBack(0)
			stack.PushBack(0)
		}
	}
	return stack.Len() == 0
}

func main() {
	str := "1,#,#,9"
	fmt.Println(isValidSerialization(str))
}
