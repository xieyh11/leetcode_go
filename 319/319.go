package main

import (
	"fmt"
	"math"
)

func bulbSwitch(n int) int {
	res := int(math.Sqrt(float64(n+1))) - 1
	if (res+2)*res < n {
		return res + 1
	}
	return res
}

func main() {
	res := make(map[int]int)
	for i := 1; i < 1500; i++ {
		res[bulbSwitch(i)]++
	}
	fmt.Println(res)
}
