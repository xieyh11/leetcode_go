package main

import (
	"fmt"
)

func singleNumber(nums []int) int {
	sum := make([]int, 64)
	for _, v := range nums {
		for i := 63; i >= 0; i-- {
			sum[i] += (v & 1)
			sum[i] %= 3
			v >>= 1
		}
	}
	res := 0
	for i := 0; i < 64; i++ {
		res <<= 1
		res |= sum[i]
	}
	return res
}

func main() {
	nums := []int{0, 1, 0, 1, 0, 1, 99}
	fmt.Println(singleNumber(nums))
}
