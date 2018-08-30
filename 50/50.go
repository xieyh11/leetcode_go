package main

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return x * x
	}
	if n == -1 {
		return 1 / x
	}
	if n == -2 {
		return 1 / (x * x)
	}
	half := n / 2
	remain := n - 2*half
	tmp := myPow(x, half)
	res := tmp * tmp
	if remain == 1 {
		res *= x
	}
	if remain == -1 {
		res /= x
	}
	return res
}

func main() {
	x := 2.0
	n := -3
	fmt.Println(myPow(x, n))
}
