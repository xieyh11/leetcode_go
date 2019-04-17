package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var reg, _ = regexp.Compile("^10*$")

func isPowerOfThree(n int) bool {
	base3 := strconv.FormatInt(int64(n), 3)
	return reg.MatchString(base3)
}

func main() {
	fmt.Println(isPowerOfThree(45))
}
