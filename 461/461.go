package main

import (
	"fmt"
)

func countOne(a int) int {
	count := 0
	for a != 0 {
		a &= a - 1
		count++
	}
	return count
}

func hammingDistance(x int, y int) int {
	return countOne(x ^ y)
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
