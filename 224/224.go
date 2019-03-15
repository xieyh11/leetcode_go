package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func opIndex(b byte) int {
	res := -1
	switch b {
	case '+':
		res = 0
	case '-':
		res = 1
	case '*':
		res = 2
	case '/':
		res = 3
	case '(':
		res = 4
	case ')':
		res = 5
	case '#':
		res = 6
	}
	return res
}

func suffixExpress(s string) []string {
	res := make([]string, 0)
	charStack := list.New()
	subS := ""
	charStack.PushBack(byte('#'))
	s += "#"
	opPriority := [7][7]int{{1, 1, -1, -1, -1, 1, 1}, {1, 1, -1, -1, -1, 1, 1}, {1, 1, 1, 1, -1, 1, 1}, {1, 1, 1, 1, -1, 1, 1}, {-1, -1, -1, -1, -1, 0, 2}, {1, 1, 1, 1, 2, 1, 1}, {-1, -1, -1, -1, -1, 2, 0}} //1 represnet first op is prior, -1 represent second is prior, 0 represent no priority, 2 represent error
	for i := range s {
		second := opIndex(s[i])
		if second >= 0 {
			if len(subS) > 0 {
				res = append(res, subS)
			}
			for true {
				firstOp := charStack.Back().Value.(byte)
				first := opIndex(firstOp)
				isBreak := false
				switch opPriority[first][second] {
				case 1:
					res = append(res, string(firstOp))
					charStack.Remove(charStack.Back())
				case -1:
					charStack.PushBack(s[i])
					isBreak = true
				case 0:
					if firstOp == '(' && s[i] == ')' {
						charStack.Remove(charStack.Back())
					}
					isBreak = true
				default:
					isBreak = true
				}
				if isBreak {
					break
				}
			}
			subS = ""
		} else {
			subS += string(s[i])
		}
	}
	return res
}

func removeEmpty(s string) string {
	res := ""
	for i := range s {
		if s[i] != ' ' {
			res += string(s[i])
		}
	}
	return res
}

func calculate(s string) int {
	s = removeEmpty(s)
	suffix := suffixExpress(s)
	datas := list.New()
	for _, ele := range suffix {
		if len(ele) == 1 && opIndex(ele[0]) >= 0 {
			data2 := datas.Remove(datas.Back()).(int)
			data1 := datas.Remove(datas.Back()).(int)
			temp := 0
			switch ele[0] {
			case '+':
				temp = data1 + data2
			case '-':
				temp = data1 - data2
			case '*':
				temp = data1 * data2
			case '/':
				temp = data1 / data2
			}
			datas.PushBack(temp)
		} else {
			data, _ := strconv.ParseInt(ele, 10, 64)
			datas.PushBack(int(data))
		}
	}
	return datas.Back().Value.(int)
}

func main() {
	s := "1 + 1"
	fmt.Println(calculate(s))
}
