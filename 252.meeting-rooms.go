package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=252 lang=golang
 *
 * [252] Meeting Rooms
 */
func canAttendMeetings(intervals [][]int) bool {
	x := make([][]int, 0)
	for i := range intervals {
		x = append(x, intervals[i])
	}
	sort.Slice(x, func(i, j int) bool {
		return x[i][0] < x[j][0]
	})
	for i := 1; i < len(x); i++ {
		if x[i][0] < x[i-1][1] {
			return false
		}
	}
	return true
}
