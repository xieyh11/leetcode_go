package main

import (
	"fmt"
)

func split(num []int, bitNum int) ([]int, []int) {
	if len(num) == 0 {
		return num, num
	}
	left := 0
	right := len(num) - 1
	for left <= right {
		if num[left]&bitNum == 0 {
			left++
		} else if num[right]&bitNum != 0 {
			right--
		} else {
			num[left], num[right] = num[right], num[left]
			left++
			right--
		}
	}
	return num[:left], num[left:]
}

func recSolOne(num []int, currentBit int) int {
	if len(num) <= 1 {
		return 0
	}
	for currentBit >= 0 {
		bitNum := 1 << uint(currentBit)
		num0, num1 := split(num, bitNum)
		if len(num0) == 0 || len(num1) == 0 {
			currentBit--
		} else {
			return recSol(num0, num1, currentBit-1)
		}
	}
	return num[0] ^ num[1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recSol(left []int, right []int, currentBit int) int {
	if len(left) == 0 {
		return recSolOne(right, currentBit)
	}
	if len(right) == 0 {
		return recSolOne(left, currentBit)
	}

	for currentBit >= 0 {
		bitNum := 1 << uint(currentBit)
		left0, left1 := split(left, bitNum)
		right0, right1 := split(right, bitNum)
		if len(left0) == 0 || len(left1) == 0 {
			if len(right0) == 0 || len(right1) == 0 {
				currentBit--
			} else {
				first := recSol(left, right0, currentBit-1)
				second := recSol(left, right1, currentBit-1)
				return getMax(first, second)
			}
		} else {
			if len(right0) == 0 || len(right1) == 0 {
				first := recSol(left0, right, currentBit-1)
				second := recSol(left1, right, currentBit-1)
				return getMax(first, second)
			} else {
				first := recSol(left0, right1, currentBit-1)
				second := recSol(left1, right0, currentBit-1)
				return getMax(first, second)
			}
		}
	}
	return left[0] ^ right[0]
}

func findMaximumXOR(nums []int) int {
	return recSolOne(nums, 30)
}

func main() {
	nums := []int{3, 10, 5, 25, 2, 8}
	fmt.Println(findMaximumXOR(nums))
}
