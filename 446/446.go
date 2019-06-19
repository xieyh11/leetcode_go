package main

import (
	"fmt"
)

func numberOfArithmeticSlices(A []int) int {
	if len(A) <= 2 {
		return 0
	}

	count := make([]map[int]int, len(A))
	for i := range count {
		count[i] = make(map[int]int)
	}
	res := 0
	for i := 1; i < len(A); i++ {
		for j := 0; j < i; j++ {
			diff := A[i] - A[j]
			sum := 0
			if v, ok := count[j][diff]; ok {
				sum = v
			}
			count[i][diff] += sum + 1
			res += sum
		}
	}
	return res
}

func main() {
	A := []int{2, 4, 6, 8, 10}
	fmt.Println(numberOfArithmeticSlices(A))

	A = []int{1, 1, 1, 1, 1}
	fmt.Println(numberOfArithmeticSlices(A))
}
