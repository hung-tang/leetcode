package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=252 lang=golang
 *
 * [252] Meeting Rooms
 */
func canAttendMeetings(intervals [][]int) bool {
	x := make(Matrix, 0)
	for i := range intervals {
		x = append(x, intervals[i])
	}
	sort.Sort(x)
	fmt.Println(x)
	for i := 1; i < len(x); i++ {
		if x[i][0] < x[i-1][1] {
			return false
		}
	}
	return true
}

type Matrix [][]int

func (a Matrix) Len() int {
	return len(a)
}

func (a Matrix) Less(i, j int) bool {
	return a[i][0] < a[j][0]
}

func (a Matrix) Swap(i, j int) {
	a[i][0], a[j][0], a[i][1], a[j][1] = a[j][0], a[i][0], a[j][1], a[i][1]
}
