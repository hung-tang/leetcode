package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	if intervals == nil || len(intervals) == 0 {
		return res
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] > prev[1] {
			res = append(res, prev)
			prev = cur
		} else {
			prev[1] = max(prev[1], cur[1])
		}
	}
	res = append(res, prev)
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

/* func main() {
}
*/
