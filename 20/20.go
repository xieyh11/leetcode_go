package main

import (
	"fmt"
)

// An integer stack
type IntStack struct {
	Stack []int
	Len   int
}

// Push to the top of the stack
func (a *IntStack) Push(val int) {
	if a.Len+1 <= len(a.Stack) {
		a.Stack[a.Len] = val
	} else {
		a.Stack = append(a.Stack, val)
	}
	a.Len++
}

// Pop from the top of the stack
func (a *IntStack) Pop() (int, bool) {
	if a.Len <= 0 {
		return 0, false
	}
	a.Len--
	return a.Stack[a.Len], true
}

func typeOfBracket(c byte) (int, bool) {
	switch c {
	case '(', ')':
		return 0, c == '('
	case '{', '}':
		return 1, c == '{'
	case '[', ']':
		return 2, c == '['
	}
	return -1, false
}

func isValid(s string) bool {
	stack := IntStack{}
	for _, c := range s {
		typeInt, isLeft := typeOfBracket(byte(c))
		if isLeft {
			stack.Push(typeInt)
		} else {
			tempInt, tempOk := stack.Pop()
			if tempOk {
				if tempInt != typeInt {
					return false
				}
			} else {
				return false
			}
		}
	}
	if stack.Len > 0 {
		return false
	} else {
		return true
	}
}

func main() {
	str := "()[]{}"

	fmt.Println(isValid(str))
}
