package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxRotateFunction(A []int) int {
	if len(A) <= 1 {
		return 0
	}

	last := len(A) - 1
	fRes := 0
	sum := 0
	for i, v := range A {
		fRes += i * v
		sum += v
	}
	maxR := fRes
	nMinus := len(A) - 1
	for last > 0 {
		fRes -= nMinus * A[last]
		fRes -= A[last]
		fRes += sum
		maxR = getMax(maxR, fRes)
		last--
	}
	return maxR
}

func main() {
	A := []int{4, 3, 2, 6}
	fmt.Println(maxRotateFunction(A))
}
