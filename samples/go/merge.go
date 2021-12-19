package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
	result := [][]int{}
	curLeftmost := intervals[0][0]
	curRightmost := intervals[0][1]
	for _, interval := range intervals {
		if interval[0] > curRightmost {
			result = append(result, []int{curLeftmost, curRightmost})
			curLeftmost = interval[0]
			curRightmost = interval[1]
			continue
		}
		if interval[1] > curRightmost {
			curRightmost = interval[1]
		}
	}
	result = append(result, []int{curLeftmost, curRightmost})
	return result
}
