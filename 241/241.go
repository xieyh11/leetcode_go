package main

import (
	"fmt"
	"strconv"
)

func isOp(c byte) bool {
	return c == '+' || c == '-' || c == '*'
}

func splitOp(input string) ([]int, []byte) {
	nums := make([]int, 0)
	ops := make([]byte, 0)

	for i := 0; i < len(input); {
		if isOp(input[i]) {
			ops = append(ops, input[i])
			i++
		} else {
			j := i + 1
			for ; j < len(input) && !isOp(input[j]); j++ {
			}
			temp, _ := strconv.ParseInt(input[i:j], 10, 64)
			nums = append(nums, int(temp))
			i = j
		}
	}

	return nums, ops
}

func opCompute(first int, op byte, second int) int {
	switch op {
	case '+':
		return first + second
	case '-':
		return first - second
	case '*':
		return first * second
	case '/':
		return first / second
	}
	return 0
}

func compute(nums []int, ops []byte) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) == 1 {
		return []int{nums[0]}
	}
	if len(nums) == 2 {
		return []int{opCompute(nums[0], ops[0], nums[1])}
	}
	res := make([]int, 0)
	for i := range ops {
		formerNums := nums[:i+1]
		formerOps := ops[:i]

		latterNums := nums[i+1:]
		latterOps := ops[i+1:]

		formerRes := compute(formerNums, formerOps)
		latterRes := compute(latterNums, latterOps)
		for _, former := range formerRes {
			for _, latter := range latterRes {
				res = append(res, opCompute(former, ops[i], latter))
			}
		}
	}
	return res
}

func diffWaysToCompute(input string) []int {
	nums, ops := splitOp(input)
	return compute(nums, ops)
}

func main() {
	input := "2*3-4*5"
	fmt.Println(diffWaysToCompute(input))
}
