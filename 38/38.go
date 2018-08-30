package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	switch n {
	case 1:
		return "1"
	case 2:
		return "11"
	case 3:
		return "21"
	case 4:
		return "1211"
	case 5:
		return "111221"
	default:
		temp := countAndSay(n - 1)
		current := temp[0]
		currentCount := 0
		res := ""
		for i := range temp {
			if temp[i] == current {
				currentCount++
			} else {
				res += strconv.Itoa(currentCount) + string(current)
				current = temp[i]
				currentCount = 1
			}
		}
		res += strconv.Itoa(currentCount) + string(current)
		return res
	}
}

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(countAndSay(i))
	}
}
