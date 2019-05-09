package main

import (
	"fmt"
)

func oneBitSum(a, b, c int) (int, int) {
	count := 0
	if a == 1 {
		count++
	}
	if b == 1 {
		count++
	}
	if c == 1 {
		count++
	}
	switch count {
	case 3:
		return 1, 1
	case 2:
		return 0, 1
	case 1:
		return 1, 0
	default:
		return 0, 0
	}
}

func getSum(a int, b int) int {
	res := 0
	carry := 0
	for i := uint(0); i < 64; i++ {
		tempR, tempC := oneBitSum(a&1, b&1, carry)
		carry = tempC
		res |= (tempR << i)
		a >>= 1
		b >>= 1
	}
	return res
}

func main() {
	fmt.Println(getSum(1, 2))
	fmt.Println(getSum(-1, 2))
	fmt.Println(getSum(-1, -2))
}
