package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	numStack := list.New()
	for _, token := range tokens {
		switch token {
		case "+":
			numb := numStack.Remove(numStack.Back()).(int)
			numa := numStack.Remove(numStack.Back()).(int)
			numStack.PushBack(numa + numb)
		case "-":
			numb := numStack.Remove(numStack.Back()).(int)
			numa := numStack.Remove(numStack.Back()).(int)
			numStack.PushBack(numa - numb)
		case "*":
			numb := numStack.Remove(numStack.Back()).(int)
			numa := numStack.Remove(numStack.Back()).(int)
			numStack.PushBack(numa * numb)
		case "/":
			numb := numStack.Remove(numStack.Back()).(int)
			numa := numStack.Remove(numStack.Back()).(int)
			numStack.PushBack(numa / numb)
		default:
			if num, err := strconv.ParseInt(token, 10, 64); err == nil {
				numStack.PushBack(int(num))
			}
		}
	}
	return numStack.Back().Value.(int)
}

func main() {
	tokens := []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}
	fmt.Println(evalRPN(tokens))
}
