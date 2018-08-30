package main

import (
	"fmt"
)

type SudoEle struct {
	Pair     [2]int
	Possible map[byte]bool
	size     int
}

func NewSudoEle(pair [2]int) *SudoEle {
	res := new(SudoEle)
	res.Pair = pair
	return res
}

func (a *SudoEle) AddPossible(ks []byte) {
	for _, v := range ks {
		if _, ok := a.Possible[v]; !ok {
			a.Possible[v] = true
			a.size++
		}
	}
}

func (a *SudoEle) RemovePossible(ks []byte) {
	for _, v := range ks {
		if _, ok := a.Possible[v]; ok {
			delete(a.Possible, v)
			a.size--
		}
	}
}

func (a *SudoEle) Use(k byte) bool {
	if v, ok := a.Possible[k]; ok {
		if v {
			a.Possible[k] = false
			a.size--
		}
		return v
	} else {
		return false
	}
}

func (a *SudoEle) Recover(k byte) bool {
	if v, ok := a.Possible[k]; ok {
		if !v {
			a.Possible[k] = true
			a.size++
		}
		return !v
	} else {
		return false
	}
}
