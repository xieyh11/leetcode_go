package main

import (
	"fmt"
)

func grayCode(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := make([]int, 2)
	res[0] = 0
	res[1] = 1
	for i := uint(1); i < uint(n); i++ {
		j := len(res) - 1
		tmp := 1 << i
		for ; j >= 0; j-- {
			res = append(res, res[j]|tmp)
		}
	}
	return res
}

func main() {
	fmt.Println(grayCode(2))
}
