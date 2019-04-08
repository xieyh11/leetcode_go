package main

import (
	"fmt"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func nthSuperUglyNumber(n int, primes []int) int {
	ugly := make([]int, 1)
	ugly[0] = 1
	tn := make([]int, len(primes))

	for n > len(ugly) {
		for i := range primes {
			for ugly[tn[i]]*primes[i] <= ugly[len(ugly)-1] {
				tn[i]++
			}
		}
		min := ugly[tn[0]] * primes[0]
		for i := range primes {
			min = getMin(min, ugly[tn[i]]*primes[i])
		}
		ugly = append(ugly, min)
	}
	return ugly[n-1]
}

func main() {
	primes := []int{2, 7, 13, 19}
	n := 12
	fmt.Println(nthSuperUglyNumber(n, primes))
}
