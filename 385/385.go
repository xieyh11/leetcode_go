package main

import (
	"fmt"
	"strconv"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val         int
	isInteger   bool
	integerList []*NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return n.isInteger
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return n.val
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.val, n.isInteger = value, true
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {
	n.isInteger = false
	n.integerList = append(n.integerList, &elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return n.integerList
}

func deserialize(s string) *NestedInteger {
	res := &NestedInteger{}
	if len(s) == 0 {
		return res
	}
	if s[0] == '[' {
		start := 1
		level := 0
		for start < len(s)-1 {
			end := start
			for end < len(s)-1 && (s[end] != ',' || level > 0) {
				if s[end] == '[' {
					level++
				} else if s[end] == ']' {
					level--
				}
				end++
			}
			res.Add(*deserialize(s[start:end]))
			start = end + 1
		}
	} else {
		num, err := strconv.ParseInt(s, 10, 64)
		if err == nil {
			res.SetInteger(int(num))
		}
	}
	return res
}

func main() {
	s := "[123,[456,[789]]]"
	fmt.Println(deserialize(s))
	s = "342"
	fmt.Println(deserialize(s))
}
