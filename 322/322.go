package main

import (
	"fmt"
	"sort"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func recSol(coins []int, amount int, store map[[2]int]int) int {
	last := len(coins) - 1
	if last == 0 {
		if amount%coins[0] == 0 {
			return amount / coins[0]
		}
		return -1
	}
	array := [2]int{len(coins), amount}
	if v, ok := store[array]; ok {
		return v
	}
	maxPossible := amount / coins[last]
	maxAmount := maxPossible * coins[last]
	res := -1
	for k := maxPossible; k >= 0; k-- {
		tempRes := recSol(coins[:last], amount-maxAmount, store)

		if tempRes != -1 {
			if res == -1 {
				res = k + tempRes
			} else {
				res = getMin(res, k+tempRes)
			}
		}

		maxAmount -= coins[last]
	}
	store[array] = res
	return res
}

func otherSol(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	count := 0
	found := false
	possibleAmount := map[int]bool{0: true}
	nextLevel := make(map[int]bool)
	for !found && len(possibleAmount) > 0 {
		count++
		for k := range possibleAmount {
			for _, coin := range coins {
				if coin+k == amount {
					found = true
					break
				} else if coin+k < amount {
					nextLevel[coin+k] = true
				}
			}
			if found {
				break
			}
		}
		possibleAmount = nextLevel
		nextLevel = make(map[int]bool)
	}
	if found {
		return count
	}
	return -1
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 {
		if amount == 0 {
			return 0
		}
		return -1
	}
	sort.Ints(coins)
	store := make(map[[2]int]int)
	return recSol(coins, amount, store)
}

func main() {
	coins := []int{86, 210, 29, 22, 402, 140, 16, 466}
	amount := 3219
	fmt.Println(coinChange(coins, amount))
	fmt.Println(otherSol(coins, amount))
}
