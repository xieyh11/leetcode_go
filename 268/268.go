package main

import (
	"fmt"
)

func missingNumber(nums []int) int {
	n := len(nums)
	all := 0
	for i := 1; i <= n; i++ {
		all ^= i
	}
	remain := 0
	for _, num := range nums {
		remain ^= num
	}
	return all ^ remain
}

func main() {
	num := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
	fmt.Println(missingNumber(num))
}
