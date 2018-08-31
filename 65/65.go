package main

import (
	"fmt"
	"regexp"
)

var pattern, _ = regexp.Compile(`^\s*[-+]?(?:\d+|\d+\.\d*|\d*\.\d+)(?:e[+-]?\d+)?\s*$`)

func isNumber(s string) bool {
	return pattern.MatchString(s)
}

func main() {
	str := "2e10"
	fmt.Println(isNumber(str))
}
