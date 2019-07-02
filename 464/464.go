package main

import (
	"fmt"
)

func secondTurn(maxChoosableInteger, desiredTotal, state int, store map[int]bool, sum int) bool {
	if v, ok := store[state]; ok {
		return v
	}
	res := true
	for i := 0; res && i < maxChoosableInteger; i++ {
		temp := 1 << uint(i)
		if temp&state == 0 {
			if sum+i+1 >= desiredTotal {
				res = false
			} else {
				newState := temp | state
				res = firstTurn(maxChoosableInteger, desiredTotal, newState, store, sum+i+1)
			}
		}
	}
	store[state] = res
	return res
}

func firstTurn(maxChoosableInteger, desiredTotal, state int, store map[int]bool, sum int) bool {
	if v, ok := store[state]; ok {
		return v
	}
	res := false
	for i := 0; !res && i < maxChoosableInteger; i++ {
		temp := 1 << uint(i)
		if temp&state == 0 {
			if sum+i+1 >= desiredTotal {
				res = true
			} else {
				newState := temp | state
				res = secondTurn(maxChoosableInteger, desiredTotal, newState, store, sum+i+1)
			}
		}
	}
	store[state] = res
	return res
}

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if (maxChoosableInteger+1)*maxChoosableInteger/2 < desiredTotal {
		return false
	}
	store := make(map[int]bool)
	state := 0
	sum := 0
	return firstTurn(maxChoosableInteger, desiredTotal, state, store, sum)
}

func main() {
	fmt.Println(canIWin(10, 11))
	fmt.Println(canIWin(15, 40))
	fmt.Println(canIWin(5, 14))
}
