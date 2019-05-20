package main

import (
	"fmt"
	"sort"
)

type SliceArray [][]int

func (this SliceArray) Less(i, j int) bool {
	if this[i][0] == this[j][0] {
		return this[i][1] < this[j][1]
	} else {
		return this[i][0] > this[j][0]
	}
}

func (this SliceArray) Len() int {
	return len(this)
}

func (this SliceArray) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func reconstructQueue(people [][]int) [][]int {
	if len(people) <= 1 {
		return people
	}
	sort.Sort(SliceArray(people))
	res := [][]int{}
	for i := range people {
		now := people[i]
		temp := append([][]int{now}, res[now[1]:]...)
		res = append(res[:now[1]], temp...)
	}
	return res
}

func main() {
	people := [][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}
	fmt.Println(reconstructQueue(people))
}
