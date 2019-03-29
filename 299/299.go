package main

import (
	"fmt"
	"strconv"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getHint(secret string, guess string) string {
	bulls := 0
	secretOther := [10]int{0}
	guessOther := [10]int{0}
	sLen := len(secret)
	for i := sLen - 1; i >= 0; i-- {
		if secret[i] == guess[i] {
			bulls++
		} else {
			secretOther[secret[i]-'0']++
			guessOther[guess[i]-'0']++
		}
	}
	cows := 0
	for i := 0; i < 10; i++ {
		cows += min(secretOther[i], guessOther[i])
	}
	return strconv.FormatInt(int64(bulls), 10) + "A" + strconv.FormatInt(int64(cows), 10) + "B"
}

func main() {
	secret := "1123"
	guess := "0111"
	fmt.Println(getHint(secret, guess))
}
