package main

import (
	"fmt"
)

func fourSumCount(A []int, B []int, C []int, D []int) int {
	cdSum := make(map[int]int)
	for i := range C {
		for j := range D {
			cdSum[C[i]+D[j]]++
		}
	}

	res := 0
	for i := range A {
		for j := range B {
			sum := A[i] + B[j]
			if v, ok := cdSum[-sum]; ok {
				res += v
			}
		}
	}
	return res
}

func main() {
	A := []int{1, 2}
	B := []int{-2, -1}
	C := []int{-1, 2}
	D := []int{0, 2}
	fmt.Println(fourSumCount(A, B, C, D))
}
