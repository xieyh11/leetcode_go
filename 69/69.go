package main

import (
	"fmt"
)

func recurSqrt(x, left, right int) int {
	if right-left == 1 {
		return left
	}
	mid := (left + right) / 2
	multi := mid * mid
	if multi == x {
		return mid
	} else if multi < x {
		return recurSqrt(x, mid, right)
	} else {
		return recurSqrt(x, left, mid)
	}
}

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	return recurSqrt(x, 1, x)
}

func main() {
	fmt.Println(mySqrt(8))
}
