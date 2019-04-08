package main

import (
	"fmt"
)

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(root int, edges [][]int, searched []bool, height []int) (int, []int) {
	minHeight := height[root]
	minRes := []int{root}
	saveHeight := height[root]
	for _, node := range edges[root] {
		if !searched[node] {
			searched[node] = true
			saveNodeHeight := height[node]
			if saveNodeHeight+1 == saveHeight {
				newRootHeight := 1
				for _, nodeOther := range edges[root] {
					if nodeOther != node {
						newRootHeight = getMax(newRootHeight, height[nodeOther]+1)
					}
				}
				height[root] = newRootHeight
			}
			height[node] = getMax(height[node], height[root]+1)
			oneHeightRes, oneNodesRes := dfs(node, edges, searched, height)
			if minHeight == oneHeightRes {
				minRes = append(minRes, oneNodesRes...)
			} else if minHeight > oneHeightRes {
				minHeight = oneHeightRes
				minRes = oneNodesRes
			}
			height[root], height[node] = saveHeight, saveNodeHeight
		}
	}
	return minHeight, minRes
}

func dfsInitial(nowAt int, edges [][]int, searched []bool, height []int) {
	height[nowAt] = 1
	for _, node := range edges[nowAt] {
		if !searched[node] {
			searched[node] = true
			dfsInitial(node, edges, searched, height)
			height[nowAt] = getMax(height[nowAt], height[node]+1)
		}
	}
}

func initialHeight(n int, edges [][]int) []int {
	searched := make([]bool, n)
	height := make([]int, n)
	searched[0] = true
	dfsInitial(0, edges, searched, height)
	return height
}

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	adjacent := make([][]int, n)
	for _, pair := range edges {
		adjacent[pair[0]] = append(adjacent[pair[0]], pair[1])
		adjacent[pair[1]] = append(adjacent[pair[1]], pair[0])
	}

	searched := make([]bool, n)
	height := initialHeight(n, adjacent)
	searched[0] = true
	_, res := dfs(0, adjacent, searched, height)
	return res
}

func main() {
	n := 0
	edges := [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}
	fmt.Println(findMinHeightTrees(n, edges))
}
