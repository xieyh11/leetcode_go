package main

import (
	"fmt"
)

type SearchValues struct {
	To    int
	Value float64
}

func bfs(s, t string, nodes []string, nodesIndex map[string]int, edges []map[int]float64) float64 {
	if _, ok := nodesIndex[s]; !ok {
		return -1
	}
	if _, ok := nodesIndex[t]; !ok {
		return -1
	}

	if s == t {
		return 1
	}

	startIdx := nodesIndex[s]
	endIdx := nodesIndex[t]
	search := []SearchValues{SearchValues{startIdx, 1}}
	searched := make([]bool, len(nodes))
	searched[startIdx] = true

	found := false
	res := float64(-1)
	for len(search) > 0 {
		secondLevel := []SearchValues{}
		for _, nowAt := range search {
			for to, v := range edges[nowAt.To] {
				if !searched[to] {
					searched[to] = true
					toValue := SearchValues{to, nowAt.Value * v}
					secondLevel = append(secondLevel, toValue)
					if to == endIdx {
						found = true
						res = toValue.Value
					}
				}
				if found {
					break
				}
			}
		}
		for _, value := range secondLevel {
			if _, ok := edges[startIdx][value.To]; !ok {
				edges[startIdx][value.To] = value.Value
			}

			if _, ok := edges[value.To][startIdx]; value.Value != 0 && !ok {
				edges[value.To][startIdx] = 1 / value.Value
			}
		}
		if found {
			break
		}
		search = secondLevel
	}
	return res
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	edges := []map[int]float64{}
	nodes := []string{}
	nodesIndex := make(map[string]int)
	for i := range equations {
		first := equations[i][0]
		second := equations[i][1]
		if _, ok := nodesIndex[first]; !ok {
			nodes = append(nodes, first)
			nodesIndex[first] = len(nodes) - 1
			edges = append(edges, make(map[int]float64))
		}
		if _, ok := nodesIndex[second]; !ok {
			nodes = append(nodes, second)
			nodesIndex[second] = len(nodes) - 1
			edges = append(edges, make(map[int]float64))
		}
		firstIdx := nodesIndex[first]
		secondIdx := nodesIndex[second]
		edges[firstIdx][secondIdx] = values[i]
		if values[i] != 0 {
			edges[secondIdx][firstIdx] = 1 / values[i]
		}
	}
	res := []float64{}
	for i := range queries {
		res = append(res, bfs(queries[i][0], queries[i][1], nodes, nodesIndex, edges))
	}
	return res
}

func main() {
	eqs := [][]string{{"a", "b"}, {"b", "c"}}
	values := []float64{2, 3}
	queries := [][]string{{"a", "c"}, {"c", "a"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}
	fmt.Println(calcEquation(eqs, values, queries))

	eqs = [][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}
	values = []float64{3.0, 4.0, 5.0, 6.0}
	queries = [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}
	fmt.Println(calcEquation(eqs, values, queries))
}
