package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	maxPro := 0
	formerMin := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-formerMin > maxPro {
			maxPro = prices[i] - formerMin
		}
		if prices[i] < formerMin {
			formerMin = prices[i]
		}
	}
	return maxPro
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
