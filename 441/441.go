package main

import (
	"fmt"
	"math"
)

func arrangeCoins(n int) int {
	if n <= 0 {
		return 0
	}
	num := int((math.Sqrt(8*float64(n)+1) - 1) / 2)
	target := num * (num + 1) / 2
	targetMore := (num + 1) * (num + 2) / 2
	if target <= n && targetMore > n {
		return num
	} else {
		return num + 1
	}
}

func main() {
	fmt.Println(arrangeCoins(5))
	fmt.Println(arrangeCoins(8))
	fmt.Println(arrangeCoins(10))
}
