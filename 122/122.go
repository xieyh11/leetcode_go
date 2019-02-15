package main

import (
	"fmt"
)

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	res := make([]int, len(prices))
	res[len(res)-1] = 0
	for i := len(res) - 2; i >= 0; i-- {
		res[i] = 0
		formerMin := prices[i]
		for j := i + 1; j < len(prices); j++ {
			if formerMin > prices[j] {
				formerMin = prices[j]
			} else {
				if j < len(prices)-1 {
					if res[i] < res[j+1]+prices[j]-formerMin {
						res[i] = res[j+1] + prices[j] - formerMin
					}
				} else {
					if res[i] < prices[j]-formerMin {
						res[i] = prices[j] - formerMin
					}
				}
			}
		}
	}
	return res[0]
}

func main() {
	prices := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))
}
