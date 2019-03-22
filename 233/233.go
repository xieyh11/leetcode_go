package main

import (
	"fmt"
)

func countDigitOne(n int) int {
	res := 0
	interval := 10
	lower := 1
	upper := 1
	for n >= lower {
		res += interval / 10 * (n / interval)
		remain := n % interval
		if remain >= lower && remain <= upper {
			res += remain - lower + 1
		} else if remain > upper {
			res += interval / 10
		}
		interval *= 10
		lower *= 10
		upper *= 10
		upper += 9
	}
	return res
}

func main() {
	fmt.Println(countDigitOne(102))
}
