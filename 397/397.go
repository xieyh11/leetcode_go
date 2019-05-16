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

func toOdd(n int) (int, int) {
	times := 0
	for n&1 == 0 {
		n /= 2
		times++
	}
	return n, times
}

func recSol(n int, store map[int]int) int {
	if n == 1 {
		return 0
	}
	if v, ok := store[n]; ok {
		return v
	}
	oddMore, addonMore := toOdd(n + 1)
	res := recSol(oddMore, store) + addonMore + 1
	oddLess, addonLess := toOdd(n - 1)
	res = getMin(res, recSol(oddLess, store)+addonLess+1)
	store[n] = res
	return res
}

func integerReplacement(n int) int {
	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}
	newN, times := toOdd(n)
	store := make(map[int]int)
	return recSol(newN, store) + times
}

func main() {
	fmt.Println(integerReplacement(8))
	fmt.Println(integerReplacement(7))
	fmt.Println(integerReplacement(200000000))
	fmt.Println(integerReplacement(2147483647))
}
