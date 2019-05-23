package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 2; i < n; i += 3 {
		res[i] = "Fizz"
	}
	for i := 4; i < n; i += 5 {
		if len(res[i]) == 0 {
			res[i] = "Buzz"
		} else {
			res[i] = "FizzBuzz"
		}
	}
	for i := 0; i < n; i++ {
		if len(res[i]) == 0 {
			res[i] = strconv.FormatInt(int64(i+1), 10)
		}
	}
	return res
}

func main() {
	fmt.Println(fizzBuzz(15))
}
