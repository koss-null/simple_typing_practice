package main

import "fmt"

var res [][]int

func copySlice(data []int) []int {
	ret := make([]int, 0, len(data))
	for i := range data {
		ret = append(ret, data[i])
	}
	return ret
}

func bfs(candidates []int, target int, curIdx int, prev []int) {
	if target == 0 {
		res = append(res, copySlice(prev))
		return
	}
	for i := curIdx; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			bfs(candidates, target-candidates[i], i, append(copySlice(prev), candidates[i]))
		}
	}
}

func dfs(candidates []int, target int, curIdx int, prev []int) [][]int {
	if target == 0 {
		return [][]int{copySlice(prev)}
	}
	res := make([][]int, 0)
	for i := curIdx; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			res = append(
				res,
				dfs(
					candidates,
					target-candidates[i],
					i,
					append(prev, candidates[i]),
				)...,
			)
		}
	}
	return res
}

func combinationSum(candidates []int, target int) [][]int {
	res = make([][]int, 0)
	if len(candidates) == 0 {
		return [][]int{}
	}
	// bfs(candidates, target, 0, []int{})
	// return res
	return dfs(candidates, target, 0, []int{})
}

func main() {
	fmt.Println(combinationSum([]int{2, 7, 6, 3, 5, 1}, 9))
}
