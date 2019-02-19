package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res ^= v
	}
	return res
}

func main() {
	nums := []int{2, 2, 1}
	fmt.Println(singleNumber(nums))
}
