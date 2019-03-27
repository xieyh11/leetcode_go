package main

import (
	"fmt"
)

func min(a, b, c int) int {
	if b < a {
		a = b
	}
	if a > c {
		a = c
	}
	return a
}

func nthUglyNumber(n int) int {
	ugly := make([]int, 1)
	ugly[0] = 1
	t2, t3, t5 := 0, 0, 0
	for n > len(ugly) {
		for ugly[t2]*2 <= ugly[len(ugly)-1] {
			t2++
		}
		for ugly[t3]*3 <= ugly[len(ugly)-1] {
			t3++
		}
		for ugly[t5]*5 <= ugly[len(ugly)-1] {
			t5++
		}
		ugly = append(ugly, min(ugly[t2]*2, ugly[t3]*3, ugly[t5]*5))
	}
	return ugly[n-1]
}

func main() {
	fmt.Println(nthUglyNumber(10))
}
