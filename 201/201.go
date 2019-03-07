package main

import (
	"fmt"
	"math/rand"
)

func rangeBitwiseAnd(m int, n int) int {
	res := 0
	for i := uint(0); i < 32; i++ {
		p := 1 << i
		if n/p == m/p && (n/p)%2 == 1 {
			res |= p
		}
	}
	return res
}

func main() {
	for i := 0; i < 10; i++ {
		m, n := rand.Intn(2147483647), rand.Intn(2147483647)
		if m > n {
			m, n = n, m
		}
		fmt.Println(m, n, rangeBitwiseAnd(m, n))
	}
}
