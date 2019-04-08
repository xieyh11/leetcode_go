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

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// dp[j] = prices[j] - min(prices[i]-dp[i-2])
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minNow := prices[0]
	res := [2]int{0, 0}
	for i := 1; i < len(prices); i++ {
		res[1], res[0], minNow = getMax(res[1], prices[i]-minNow), res[1], getMin(minNow, prices[i]-res[0])
	}
	return res[1]
}

func main() {
	prices := []int{1, 2, 3, 0, 2}
	fmt.Println(maxProfit(prices))
}
