package main

import (
	"fmt"
)

type Node struct {
	Str  string
	Last int
}

func suffixUpdate(num int, str string, former map[int][]Node) map[int][]Node {
	tempSuffix := make(map[int][]Node)
	for k, v := range former {
		for _, oneStr := range v {
			if oneStr.Last != 0 {
				newValue := k / oneStr.Last * (oneStr.Last*10 + num)
				tempSuffix[newValue] = append(tempSuffix[newValue], Node{oneStr.Str + str, oneStr.Last*10 + num})
			}
			tempSuffix[k*num] = append(tempSuffix[k*num], Node{oneStr.Str + "*" + str, num})
		}
	}
	return tempSuffix
}

func minusRevert(str string) string {
	strByte := []byte(str)
	for i := range strByte {
		switch strByte[i] {
		case '+':
			strByte[i] = '-'
		case '-':
			strByte[i] = '+'
		}
	}
	return string(strByte)
}

func recSol(nums []int, strs []string, target int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return []string{strs[0]}
		}
		return []string{}
	}
	suffixPossible := make(map[int][]Node)
	suffixPossible[nums[0]] = []Node{Node{strs[0], nums[0]}}
	res := make([]string, 0)
	for i := 0; i < len(nums)-1; i++ {
		if i > 0 {
			suffixPossible = suffixUpdate(nums[i], strs[i], suffixPossible)
		}
		for k, v := range suffixPossible {
			plusRes := recSol(nums[i+1:], strs[i+1:], target-k)
			if len(plusRes) > 0 {
				for _, fromer := range v {
					for _, latter := range plusRes {
						res = append(res, fromer.Str+"+"+latter)
					}
				}
			}
			minusRes := recSol(nums[i+1:], strs[i+1:], k-target)
			if len(minusRes) > 0 {
				for j := range minusRes {
					minusRes[j] = minusRevert(minusRes[j])
				}
				for _, fromer := range v {
					for _, latter := range minusRes {
						res = append(res, fromer.Str+"-"+latter)
					}
				}
			}
		}
	}
	suffixPossible = suffixUpdate(nums[len(nums)-1], strs[len(strs)-1], suffixPossible)
	if v, ok := suffixPossible[target]; ok {
		for _, oneV := range v {
			res = append(res, oneV.Str)
		}
	}
	return res
}

func addOperators(num string, target int) []string {
	nums := make([]int, len(num))
	strs := make([]string, len(num))
	for i := range num {
		nums[i] = int(num[i] - '0')
		strs[i] = string(num[i])
	}
	return recSol(nums, strs, target)
}

func main() {
	num := "123"
	target := 6
	fmt.Println(addOperators(num, target))
}
