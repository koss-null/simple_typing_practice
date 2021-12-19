package main

import (
	"fmt"
)

var res [][]int

func copySlice(data []int) []int {
	ret := make([]int, 0, len(data))
	for i := range data {
		ret = append(ret, data[i])
	}
	return ret
}

var used map[[51]int]struct{}

func getHash(data []int) (res [51]int) {
	for i := range data {
		res[data[i]]++
	}
	return
}

func bfs(candidates []int, target int, curIdx int, prev []int) {
	if target == 0 {
		copied := copySlice(prev)
		hashCopied := getHash(copied)
		if _, ok := used[hashCopied]; !ok {
			res = append(res, copied)
			used[hashCopied] = struct{}{}
		}
		return
	}
	for i := curIdx; i < len(candidates); i++ {
		if target-candidates[i] >= 0 {
			newSlice := append(copySlice(prev), candidates[i])
			hash := getHash(newSlice)
			if _, ok := used[hash]; ok {
				continue
			}
			bfs(candidates, target-candidates[i], i+1, newSlice)
			used[hash] = struct{}{}
		}
	}
}

func compare(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func combinationSum2(candidates []int, target int) [][]int {
	used = map[[51]int]struct{}{}
	if len(candidates) == 0 {
		return [][]int{}
	}

	res = make([][]int, 0)
	bfs(candidates, target, 0, []int{})
	return res
}

func main() {
	fmt.Println(combinationSum2([]int{1}, 9))
}
