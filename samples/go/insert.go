package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func insert(intervals [][]int, newInterval []int) [][]int {
	result := [][]int{}
	for i := range intervals {
		if newInterval[0] >= intervals[i][0] && intervals[i][1] >= newInterval[0] {
			curLeftmost := intervals[i][0]
			curRightmost := max(intervals[i][1], newInterval[1])
			for _, interval := range intervals[i:] {
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
		result = append(result, intervals[i])
	}

	result = [][]int{}
	var i int
	for i = range intervals {
		if newInterval[0] >= intervals[i][0] {
			result = append(result, intervals[i])
		}
		if intervals[i][1] < newInterval[0] || newInterval[1] < intervals[i][0] {
			break
		}
	}

	curLeftmost := newInterval[0]
	curRightmost := newInterval[1]
	for _, interval := range intervals[i:] {
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
