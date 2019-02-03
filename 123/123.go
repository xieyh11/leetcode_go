package main

import (
	"fmt"
)

func oneMaxProfit(prices []int) (int, int) {
	if len(prices) < 2 {
		return 0, -1
	}
	formerMin := prices[0]
	formerIdx := 0
	max := 0
	maxIdx := -1
	for i := 1; i < len(prices); i++ {
		if max < prices[i]-formerMin {
			max = prices[i] - formerMin
			maxIdx = formerIdx
		}
		if formerMin > prices[i] {
			formerMin = prices[i]
			formerIdx = i
		}
	}
	return max, maxIdx
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	res := make([][]int, 2)
	res[0] = make([]int, len(prices))
	res[1] = make([]int, len(prices))

	for i := 0; i < len(prices); {
		maxNow, startAt := oneMaxProfit(prices[i:])
		if startAt != -1 {
			for j := 0; j <= startAt; j++ {
				res[1][i+j] = maxNow
			}
			i += startAt + 1
		} else {
			break
		}
	}

	res[0][len(prices)-1] = 0
	res[0][len(prices)-2] = res[1][len(prices)-2]
	for i := len(prices) - 3; i >= 0; i-- {
		res[0][i] = res[0][i+1]
		for j := i + 1; j < len(prices)-1; j++ {
			if prices[j] > prices[i] {
				if res[0][i] < res[1][j+1]+prices[j]-prices[i] {
					res[0][i] = res[1][j+1] + prices[j] - prices[i]
				}
			}
		}
		if res[0][i] < res[1][i] {
			res[0][i] = res[1][i]
		}
	}
	return res[0][0]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit(nums))
}
