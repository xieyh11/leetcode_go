package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func recSol(i, k int, res [][]int, prices []int) {
	if res[k][i] != -1 {
		return
	}

}

type ProfitNode struct {
	Profit int
	Times  int
}

func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	if k < 1 {
		return 0
	}
	nPrices := len(prices)
	noLimits := make([]ProfitNode, nPrices)
	noLimits[nPrices-1] = ProfitNode{0, 0}
	for i := nPrices - 2; i >= 0; i-- {
		noLimits[i] = ProfitNode{0, 0}
		formerMin := prices[i]
		for j := i + 1; j < len(prices); j++ {
			if formerMin > prices[j] {
				formerMin = prices[j]
			} else {
				if j < len(prices)-1 {
					if noLimits[i].Profit < noLimits[j+1].Profit+prices[j]-formerMin {
						noLimits[i].Profit = noLimits[j+1].Profit + prices[j] - formerMin
						noLimits[i].Times = noLimits[j+1].Times + 1
					}
				} else {
					if noLimits[i].Profit < prices[j]-formerMin {
						noLimits[i] = ProfitNode{prices[j] - formerMin, 1}
					}
				}
			}
		}
	}

	if k >= noLimits[0].Times {
		return noLimits[0].Profit
	}

	res := make([][]int, 2)
	res[0] = make([]int, nPrices)
	res[1] = make([]int, nPrices)
	res[0][nPrices-1] = 0
	maxAfter := prices[nPrices-1]
	for j := nPrices - 2; j >= 0; j-- {
		res[0][j] = getMax(res[0][j+1], maxAfter-prices[j])
		if maxAfter < prices[j] {
			maxAfter = prices[j]
		}
	}
	for now := k - 2; now >= 0; now-- {
		copy(res[1], res[0])
		res[0][nPrices-1] = 0
		for i := nPrices - 2; i >= 0; i-- {
			if noLimits[i].Times > k-now {
				res[0][i] = getMax(res[0][i+1], res[1][i])
				for j := i + 1; j < nPrices; j++ {
					if prices[j] > prices[i] {
						temp := prices[j] - prices[i]
						if j < nPrices-1 {
							temp += res[1][j+1]
						}
						res[0][i] = getMax(res[0][i], temp)
					}
				}
			} else {
				res[0][i] = noLimits[i].Profit
			}
		}
	}
	return res[0][0]
}

func main() {
	prices := []int{3, 2, 6, 5, 0, 3}
	k := 2
	fmt.Println(maxProfit(k, prices))
}
