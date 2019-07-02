package main

import (
	"fmt"
	"math/rand"
)

func rand7() int {
	return rand.Intn(7) + 1
}

func rand10() int {
	choose := rand7()
	for choose == 7 {
		choose = rand7()
	}
	another := rand7()
	for another > 5 {
		another = rand7()
	}
	if choose <= 3 {
		return another
	} else {
		return another + 5
	}
}

func main() {
	fmt.Println(rand10())
}
