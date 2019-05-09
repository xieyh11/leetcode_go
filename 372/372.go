package main

import (
	"fmt"
)

func isZero(b []int) bool {
	return len(b) == 1 && b[0] == 0
}

func divide2(b []int) ([]int, bool) {
	res := []int{}
	carry := 0
	for i := 0; i < len(b); i++ {
		carry = carry*10 + b[i]
		res = append(res, carry/2)
		carry %= 2
	}
	for i := 0; i < len(res); i++ {
		if res[i] != 0 {
			res = res[i:]
			break
		}
	}
	return res, carry == 1
}

func superPow(a int, b []int) int {
	a %= 1337
	if isZero(b) {
		return 1
	}
	half, remain := divide2(b)
	halfPow := superPow(a, half)
	res := (halfPow * halfPow) % 1337
	if remain {
		res = (res * a) % 1337
	}
	return res
}

func main() {
	a := 2
	b := []int{1, 0}
	fmt.Println(superPow(a, b))

	a = 2
	b = []int{3}
	fmt.Println(superPow(a, b))
}
