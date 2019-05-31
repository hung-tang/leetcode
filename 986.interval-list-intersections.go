package main

import (
	"sort"
)

/*
 * @lc app=leetcode id=986 lang=golang
 *
 * [986] Interval List Intersections
 */
func intervalIntersection(A [][]int, B [][]int) [][]int {
	res := make([][]int, 0)
	if A == nil || len(A) == 0 || B == nil || len(B) == 0 {
		return res
	}
	intervals := make([][]int, 0)
	for _, each := range A {
		intervals = append(intervals, each)
	}
	for _, each := range B {
		intervals = append(intervals, each)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] > prev[1] {
			prev = cur
			continue
		}
		// 3 cases: (1) prev is longer, (2) cur is longer, (3) prev and cur same length
		if cur[1] == prev[1] { // case 3
			res = append(res, []int{cur[0], cur[1]})
			prev = cur
		} else if cur[1] > prev[1] { // case 2
			res = append(res, []int{cur[0], cur[0] + prev[1] - cur[0]})
			prev = cur
		} else { // case 1
			res = append(res, []int{cur[0], cur[1]})
		}
	}
	return res
}

/* func main() {
	a := [][]int{
		{0, 2},
		{5, 10},
		{13, 23},
		{24, 25},
	}
	b := [][]int{
		{1, 5},
		{8, 12},
		{15, 24},
		{25, 26},
	}
	fmt.Println(intervalIntersection(a, b))
} */
