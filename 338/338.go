package main

import (
	"fmt"
)

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	res := make([]int, num+1)
	res[0] = 0
	res[1] = 1
	for i := 2; i <= num; i++ {
		if i&1 == 0 {
			res[i] = res[i/2]
		} else {
			res[i] = res[i/2] + 1
		}
	}
	return res
}

func main() {
	fmt.Println(countBits(10))
}
